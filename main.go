package main

import (
	"go-jwt/database"
	"go-jwt/router"
)

func main() {

	database.StartDB()

	// router.StartApp().Run(":8080")
	r := router.StartApp()
	r.Run(":8080")

}
