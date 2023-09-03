package entities

import "gorm.io/gorm"

// Book is a struct that represents a book
type BookModel struct {
	gorm.Model
	ID     int
	Title  string
	Author string
}
