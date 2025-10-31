package main
import (
	. "library_system/models"
	. "library_system/services"
	. "library_system/controllers"
)

func main() {
	library_service := Library{make(map[int]Book), make(map[int]Member)}
	library_controller := LibraryController{}
	library_controller.SetLibraryService(library_service)
	library_controller.Run()
}