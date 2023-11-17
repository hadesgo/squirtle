package main

import (
	_ "squirtle/database"
	"squirtle/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":8080")
}
