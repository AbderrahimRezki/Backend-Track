package controllers
import (
	"fmt"
	. "library_system/services"
	. "library_system/models"
)

type Controller interface {
	Run()
	ShowMenu()
	MakeChoice(int) error
}

type LibraryController struct {
	library_service Library
}

func (controller *LibraryController) SetLibraryService(lib Library) {
	controller.library_service = lib
}

func (controller *LibraryController) Run() {
	var choice int
	for {
		controller.ShowMenu()
		fmt.Scanf("%d", &choice)
		controller.MakeChoice(choice)

		if choice == 8 {
			break
		}
	}
}

func (controller *LibraryController) ShowMenu() {
	menu := `
	Choose operation number:
	1 - Add Member
	2 - Add Book
	3 - Remove Book
	4 - Borrow Book
	5 - Return Book
	6 - List Availabe Books
	7 - List Borrowed Books
	8 - Exit
	`
	fmt.Println(menu)
}

func (controller *LibraryController) MakeChoice(choice int) {
	switch choice {
	case 1:
		var ID int
		var name string
		fmt.Println("Enter member details as follows: ID Name")
		fmt.Scanf("%d %s", &ID, &name)
		err := controller.library_service.AddMember(Member {
			ID: ID,
			Name: name,
			BorrowedBooks: make([]Book, 0),
		})

		if err != nil {
			fmt.Println(err.Error())
		}
	
	case 2: 
		var ID int
		var title, author string
		fmt.Println("Enter book details as follows: ID Title Author")
		fmt.Scanf("%d %s %s", &ID, &title, &author)
		err := controller.library_service.AddBook(Book {
			ID: ID,
			Title: title,
			Author: author,
			Status: Available,
		})

		if err != nil {
			fmt.Println(err.Error())
		}


	case 3:
		var ID int
		fmt.Println("Enter book ID to remove")
		fmt.Scanf("%d", &ID)
		err := controller.library_service.RemoveBook(ID)

		if err != nil {
			fmt.Println(err.Error())
		}
	
	case 4:
		var bookID, memberID int
		fmt.Println("Enter Book ID and Member ID to borrow")
		fmt.Scanf("%d %d", &bookID, &memberID)
		err := controller.library_service.BorrowBook(bookID, memberID)

		if err != nil {
			fmt.Println(err.Error())
		}

	case 5:
		var bookID, memberID int
		fmt.Println("Enter Book ID and Member ID to return")
		fmt.Scanf("%d %d", &bookID, &memberID)
		err := controller.library_service.ReturnBook(bookID, memberID)

		if err != nil {
			fmt.Println(err.Error())
		}

	case 6:
		fmt.Println("Available Books: ")
		fmt.Println(controller.library_service.ListAvailableBooks())
	
	case 7:
		var memberID int

		fmt.Println("Enter Member ID: ")
		fmt.Scanf("%d", &memberID)
		fmt.Println(controller.library_service.ListBorrowedBooks(memberID))

	case 8:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Please choose a valid operation")
	}	
}