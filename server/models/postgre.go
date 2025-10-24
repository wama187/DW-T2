package models

import (
	"time"
)

type UserModel struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookModel struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Author    string
	OwnerID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}