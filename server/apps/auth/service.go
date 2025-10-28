package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"time"

	"golang.org/x/oauth2"
	"github.com/golang-jwt/jwt/v5"
	"app/server/apps/user"
	"app/server/gen/auth"
	
)

type TokenResponse struct {
		AccessToken string
		TokenType   string
		ExpiresIn  int
	}

type authsrvc struct{
	repo user.UserRepository
	jwtSecret string
	tokenTTL time.Duration
	cfg *oauth2.Config
}

func NewAuth(repo user.UserRepository, jwtSecret string, TTL time.Duration, cfg *oauth2.Config)  *authsrvc {
	return &authsrvc{
		repo: repo,
		jwtSecret: jwtSecret,
		tokenTTL: TTL,
		cfg: cfg,
	}
}

func (s *authsrvc) GoogleLogin(ctx context.Context) (*auth.GoogleLoginResult, error) {
	url := s.cfg.AuthCodeURL("randomstate", oauth2.AccessTypeOffline)
	return &auth.GoogleLoginResult{RedirectURL: url}, nil
}

func (s *authsrvc) GoogleCallback(ctx context.Context, payload *auth.GoogleCallbackPayload) (*auth.GoogleCallbackResult, error) {
	code := payload.Code
	if code == "" {
		return nil, fmt.Errorf("missing code")
	}

	token, err := s.cfg.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("token exchange failed: %v", err)
	}

	client := s.cfg.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email   string `json:"email"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	usr, err := s.repo.FindByEmail(userInfo.Email)
	if err != nil {
		newUser := &user.User{
			Email:    userInfo.Email,
			Name:     userInfo.Name,
			Password: "oauthlogin",
		}

		if err := s.repo.Create(newUser); err != nil {
			return nil, fmt.Errorf("error al crear el usuario: %v", err)
		}
		usr = newUser
	}


	tokenStr := s.generateToken(usr.ID)

	return &auth.GoogleCallbackResult{
		Email:   userInfo.Email,
		Name:    userInfo.Name,
		Picture: &userInfo.Picture,
		Token:   tokenStr.AccessToken,
	}, nil
}

func (s *authsrvc) generateToken(userID uint) *TokenResponse {
	
	now := time.Now()
	claims := jwt.MapClaims{
		"user_id": userID,
		"iat":     now.Unix(),
		"exp":     now.Add(s.tokenTTL).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil
	}

	TTL := int(s.tokenTTL.Seconds())

	return &TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   TTL,
	}
}