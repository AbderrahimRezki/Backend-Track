# Library Management System

A simple command-line library management system that allows to manage books and members.

## Project Structure

```
.
├── go.mod
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   ├── book.go
│   └── member.go
├── services/
│   └── library_service.go
└── utils/
    └── find_by_id.go
```

## Components

### Models

#### Book
Represents a book in the library with the following fields:
- `ID`: Unique identifier
- `Title`: Book title
- `Author`: Author name
- `Status`: Either "Available" or "Borrowed"

#### Member
Represents a library member with:
- `ID`: Unique identifier
- `Name`: Member name
- `BorrowedBooks`: List of books currently borrowed

### Services

#### Library Service
Implements the core library operations:
- `AddMember`: Register a new member
- `AddBook`: Add a book to the collection
- `RemoveBook`: Remove a book from the library
- `BorrowBook`: Allow a member to borrow an available book
- `ReturnBook`: Process book returns
- `ListAvailableBooks`: Get all available books
- `ListBorrowedBooks`: Get books borrowed by a specific member

### Controllers

#### Library Controller
Handles user interaction through a menu-driven interface. Manages the flow between user input and library service operations.

### Utils

#### FindByID
Generic utility function that finds elements in a slice by their ID. Works with any type that implements the `HasID` interface, here (Member and Book).

## Running the Application

```bash
go run main.go
```