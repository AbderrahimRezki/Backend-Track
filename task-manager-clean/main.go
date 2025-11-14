package main

import (
	"task-manager-clean/delivery/routers"
)

func main() {
	router := routers.GetRouter()
	router.Run()
}
