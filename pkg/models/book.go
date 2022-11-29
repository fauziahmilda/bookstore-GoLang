// file ini lbh besar, jadi terakhir setelah util
package models

import (
	"github.com/fauziahmilda/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm: "" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CREATE FUNCTION
func (b *Book) CreateBook() *Book {
	//db help to talk to database
	db.NewRecord(b)
	//b isi something will we receive of type book
	db.Create(&b)
	return b
}

// GET ALL FUNCTION
func GetAllBooks() []Book { //using slice return list of database
	var Books []Book
	db.Find(&Books)
	return Books
}

// GET A BOOK BY ID FUNCTION
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// DELETE FUNCTION
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}

// UPDATE FUNCTION
// func UpdateBook() {
// 	var
// }
