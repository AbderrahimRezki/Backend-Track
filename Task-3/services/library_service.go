package services
import (
	"fmt"
	. "library_system/models"
	. "library_system/utils"
)

type LibraryManager interface {
	AddMember(member Member) error
	AddBook(book Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}

type Library struct {
	Books map[int]Book
	Members map[int]Member
}

func (lib Library) AddMember(member Member) error {
	if _, ok := lib.Members[member.ID]; !ok {
		lib.Members[member.ID] = member
		return nil
	}

	err := fmt.Errorf("Member %v already exists", member)
	return err
}

func (lib Library) AddBook(book Book) error {
	if _, ok := lib.Books[book.ID]; !ok {
		lib.Books[book.ID] = book
		return nil
	}

	err := fmt.Errorf("Book %v already exists", book)
	return err
}

func (lib Library) RemoveBook(bookID int) error{
	if _, ok := lib.Books[bookID]; !ok {
		err := fmt.Errorf("Book ID %v does not exist", bookID)
		return err
	}
	delete(lib.Books, bookID)
	return nil
}

func (lib Library) BorrowBook(bookID int, memberID int) error {
	if _, ok := lib.Books[bookID]; !ok {
		err := fmt.Errorf("Book ID %v does not exist", bookID)
		return err
	}

	book := lib.Books[bookID]
	if book.Status != Available {
		err := fmt.Errorf("Book %v unavailable", book)
		return err
	}

	member := lib.Members[memberID]
	book.Status = Borrowed
	newBooks := append(lib.Members[memberID].BorrowedBooks, book)
	member.BorrowedBooks = newBooks
	lib.Members[memberID] = member
	lib.Books[bookID] = book
	return nil
}

func (lib Library) ReturnBook(bookID int, memberID int) error {
	if _, ok := lib.Members[memberID]; !ok {
		err := fmt.Errorf("Member %v not found", memberID)
		return err
	}

	member := lib.Members[memberID]
	idx := FindByID(member.BorrowedBooks, bookID)
	if idx == -1 {
		err := fmt.Errorf("Member %v does not have %v", memberID, bookID)
		return err
	}

	book := lib.Books[bookID]
	book.Status = Available
	lib.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks[:idx], member.BorrowedBooks[idx+1:]...)
	return nil
}

func (lib Library) ListAvailableBooks() []Book {
	var availabeBooks []Book
	for _, book := range lib.Books {
		if book.Status == Available {
			availabeBooks = append(availabeBooks, book)
		}
	}

	return availabeBooks
}

func (lib Library) ListBorrowedBooks(memberID int) []Book {
	var borrowedBooks []Book
	for _, book := range lib.Members[memberID].BorrowedBooks {
		if book.Status == Borrowed {
			borrowedBooks = append(borrowedBooks, book)
		}
	}

	return borrowedBooks
}