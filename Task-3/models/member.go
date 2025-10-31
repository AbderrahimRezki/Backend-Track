package models
import "fmt"

type Member struct {
	ID int
	Name string
	BorrowedBooks []Book
}

func (member Member) GetID() int {
	return member.ID
}

func (member Member) String() string {
	return fmt.Sprintf("ID: %v | Name: %s | Borrowed Books: %v", member.ID, member.Name, member.BorrowedBooks)
}