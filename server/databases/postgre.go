package databases

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"app/server/models"
)

func ConnectPostgres() (*gorm.DB, error) {
	host := "postgres"
	port := 5432
	user := "user"
	password := "password"
	dbname := "db"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate() error {
	db, err := ConnectPostgres()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.UserModel{}, &models.BookModel{})
	if err != nil {
		return err
	}

	return nil
}

