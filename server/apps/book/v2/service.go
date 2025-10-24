package book_v2

import (
	"app/server/gen/books_v2"
	"app/server/apps/book"
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"goa.design/goa/v3/security"
)

type contextKey string

const userIDKey contextKey = "user_id"

type BookService struct {
	repo      book.BookRepository[book.BookMongo, string, uint]
	jwtSecret string
}

func NewService(repo book.BookRepository[book.BookMongo, string, uint], jwtSecret string) *BookService {
	return &BookService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *BookService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", t.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return ctx, fmt.Errorf("token inválido: %w", err)
	}

	if !parsedToken.Valid {
		return ctx, fmt.Errorf("token no válido")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return ctx, fmt.Errorf("claims inválidos")
	}


	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return ctx, fmt.Errorf("user_id no encontrado en el token")
	}

	userID := uint(userIDFloat)
	
	ctx = context.WithValue(ctx, userIDKey, userID)
	return ctx, nil
}

func userIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(userIDKey).(uint)
	if !ok {
		return 0, fmt.Errorf("user_id no encontrado en el contexto")
	}
	return userID, nil
}

func (s *BookService) Create(ctx context.Context, p *booksv2.BookPayloadV2) error {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return err
	}

	book := &book.BookMongo{
		Title:   p.Title,
		Author:  p.Author,
		OwnerID: userID,
	}

	return s.repo.Create(book)
}

func (s *BookService) Update(ctx context.Context, p *booksv2.UpdateBookPayloadV2) error {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return err
	}

	existingBook, err := s.repo.FindByID(&book.BookMongo{
		ID:      *p.ID,
		OwnerID: userID,
	})
	if err != nil {
		return fmt.Errorf("libro no encontrado")
	}

	if existingBook.OwnerID != userID {
		return fmt.Errorf("no tienes permiso para actualizar este libro")
	}

	existingBook.Title = p.Title
	existingBook.Author = p.Author
	return s.repo.Update(existingBook)
}

func (s *BookService) List(ctx context.Context, p *booksv2.ListPayload) ([]*booksv2.BookV2, error) {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	domainBooks, err := s.repo.FindAllByOwner(userID)
	if err != nil {
		return nil, err
	}

	result := make([]*booksv2.BookV2, len(domainBooks))
	for i, b := range domainBooks {
		result[i] = &booksv2.BookV2{
			ID:      b.ID,
			Title:   b.Title,
			Author:  b.Author,
			OwnerID: b.OwnerID,
		}
	}
	return result, nil
}

func (s *BookService) Delete(ctx context.Context, p *booksv2.DeletePayload) error {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return err
	}

	existingBook, err := s.repo.FindByID(&book.BookMongo{
		ID:      p.ID,
		OwnerID: userID,
	})
	if err != nil {
		return fmt.Errorf("libro no encontrado")
	}

	if existingBook.OwnerID != userID {
		return fmt.Errorf("no tienes permiso para eliminar este libro")
	}

	return s.repo.Delete(p.ID, userID)
}
	