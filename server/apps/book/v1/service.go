package book_v1

import (
	"app/server/gen/books"
	"app/server/apps/book"
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"goa.design/goa/v3/security"
)

type contextKey string

const userIDKey contextKey = "user_id"

type BookService struct {
	repo      book.BookRepository[book.BookPostgre, uint, uint]
	jwtSecret string
}

func NewService(repo book.BookRepository[book.BookPostgre, uint, uint], jwtSecret string) *BookService {
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

func (s *BookService) Create(ctx context.Context, p *books.BookPayloadV1) error {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return err
	}

	newBook := &book.BookPostgre{
		Title:   p.Title,
		Author:  p.Author,
		OwnerID: userID,
	}

	return s.repo.Create(newBook)
}

func (s *BookService) Update(ctx context.Context, p *books.UpdateBookPayloadV1) error {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return err
	}

	existingBook, err := s.repo.FindByID(&book.BookPostgre{
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

func (s *BookService) List(ctx context.Context, p *books.ListPayload) ([]*books.BookV1, error) {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	domainBooks, err := s.repo.FindAllByOwner(userID)
	if err != nil {
		return nil, err
	}

	result := make([]*books.BookV1, len(domainBooks))
	for i, b := range domainBooks {
		result[i] = &books.BookV1{
			ID:      b.ID,
			Title:   b.Title,
			Author:  b.Author,
			OwnerID: b.OwnerID,
		}
	}
	return result, nil
}

func (s *BookService) Delete(ctx context.Context, p *books.DeletePayload) error {
	userID, err := userIDFromContext(ctx)
	if err != nil {
		return err
	}

	existingBook, err := s.repo.FindByID(&book.BookPostgre{
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
