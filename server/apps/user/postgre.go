package user

import (
	"gorm.io/gorm"
	"app/server/models"
	"time"
)

type userRepoPostgres struct {
	db *gorm.DB
}

func NewUserRepositoryPostgres(db *gorm.DB) UserRepository {
	return &userRepoPostgres{db: db}
}

func (r *userRepoPostgres) Create(u *User) error {
	user := &models.UserModel{
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return r.db.Create(user).Error
}

func (r *userRepoPostgres) FindByEmail(email string) (*User, error) {
	var user models.UserModel
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	
	return &User{
		ID: 	  user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}