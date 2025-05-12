package models

import (
	"github.com/jinzhu/gorm"
	"github.com/krmohit2k3/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"` // gorm:"" has deleted from this
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// functions
// create book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(*b)
	return b
}

// get all books
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// get book by Id
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// delete a book
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
