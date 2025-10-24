package book_v1

import (
	"app/server/apps/book"
	"app/server/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type bookRepoPostgres struct {
	db *gorm.DB
}

func NewBookRepositoryPostgres(db *gorm.DB) *bookRepoPostgres {
	return &bookRepoPostgres{db: db}
}

func (r *bookRepoPostgres) Create(b *book.BookPostgre) error {
	if b.OwnerID == 0 {
			return fmt.Errorf("owner_id is required")
		}

	book := &models.BookModel{
		Title:     b.Title,
		Author:    b.Author,
		OwnerID:   b.OwnerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return r.db.Create(book).Error
}

func (r *bookRepoPostgres) Update(b *book.BookPostgre) error {
	book := &models.BookModel{}
	r.db.First(&book, b.ID)
	book.Title = b.Title
	book.Author = b.Author
	book.UpdatedAt = time.Now()
	return r.db.Save(book).Error
}

func (r *bookRepoPostgres) Delete(bookID uint, ownerID uint) error {
	return r.db.Where("id = ? AND owner_id = ?", bookID, ownerID).Delete(&models.BookModel{}).Error
}

func (r *bookRepoPostgres) FindByID(b *book.BookPostgre) (*book.BookPostgre, error) {
	var bookModel models.BookModel
	err := r.db.Where("id = ? AND owner_id = ?", b.ID, b.OwnerID).First(&bookModel).Error
	if err != nil {
		return nil, err
	}

	return &book.BookPostgre{
		ID:      bookModel.ID,
		Title:   bookModel.Title,
		Author:  bookModel.Author,
		OwnerID: bookModel.OwnerID,
	}, nil
}

func (r *bookRepoPostgres) FindAllByOwner(ownerID uint) ([]*book.BookPostgre, error) {
	var books []models.BookModel
	err := r.db.Where("owner_id = ?", ownerID).Find(&books).Error
	result := make([]*book.BookPostgre, len(books))
	for i, b := range books {
		result[i] = &book.BookPostgre{
			ID:      b.ID,
			Title:   b.Title,
			Author:  b.Author,
			OwnerID: b.OwnerID,
		}
	}
	return result, err
}
