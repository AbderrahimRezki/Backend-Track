package main

import (
	"fmt"
	"task-manager/router"
)

func main() {
	router := router.GetRouter()
	if err := router.Run("localhost:8080"); err != nil {
		fmt.Println("Could not start server; Encountered error: " + err.Error())
	}
}
