package databases

import (
	"fmt"
	"time"
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

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Conexión a Postgres exitosa")
			return db, nil
		}
		fmt.Printf("Intento %d: Postgres no listo, reintentando en 2s...\n", i+1)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("no se pudo conectar a la DB después de varios intentos: %w", err)
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

	fmt.Println("Migración completada")
	return nil
}
