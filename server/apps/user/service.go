package user

import (
	"errors"
	"fmt"
	"time"
	"context"
	"app/server/gen/users"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      UserRepository
	jwtSecret string
	tokenTTL  time.Duration
}

func NewUserService(repo UserRepository, jwtSecret string, tokenTTL time.Duration) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: jwtSecret,
		tokenTTL:  tokenTTL,
	}
}

func (s *UserService) Register(ctx context.Context, p *users.RegisterPayload) (*users.TokenResponse, error) {
	_, err := s.repo.FindByEmail(p.Email)
	if err == nil {
		return nil, errors.New("usuario ya existe")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error al hashear contraseña: %w", err)
	}

	user := &User{
		Name:     p.Name,
		Email:    p.Email,
		Password: string(hash),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, fmt.Errorf("error al crear usuario: %w", err)
	}

	return s.generateToken(user.ID), nil
}

func (s *UserService) Login(ctx context.Context, p *users.LoginPayload) (*users.TokenResponse, error) {
	user, err := s.repo.FindByEmail(p.Email)
	if err != nil {
		return nil, errors.New("credenciales incorrectas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password)); err != nil {
		return nil, errors.New("credenciales incorrectas")
	}

	return s.generateToken(user.ID), nil
}

func (s *UserService) generateToken(userID uint) *users.TokenResponse {
	
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

	return &users.TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   &TTL,
	}
}

func (s *UserService) ValidateToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", t.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return 0, fmt.Errorf("error al parsear token: %w", err)
	}

	if !token.Valid {
		return 0, errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("claims inválidos")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id no encontrado o tipo incorrecto")
	}

	return int(userIDFloat), nil
}

