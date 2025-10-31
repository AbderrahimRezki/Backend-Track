package models
import "fmt"

type BookStatus string
const (
	Available BookStatus = "Available"
	Borrowed = "Borrowed"
)

type Book struct {
	ID int
	Title string
	Author string
	Status BookStatus
}

func (book Book) GetID() int {
	return book.ID
}

func (book Book) String() string {
	return fmt.Sprintf("ID: %v | Title: %s | Author: %s | Status: %s", book.ID, book.Title, book.Author, book.Status)
}