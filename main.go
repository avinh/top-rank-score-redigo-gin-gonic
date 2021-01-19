package main

import (
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	handlers := handlers.ConnRedis{}
	handlers.Connect()

	router := gin.Default()
	router.GET("/add-user", handlers.AddUser)
	router.GET("/get-top", handlers.GetTop)
	router.GET("/get-user", handlers.GetUser)
	router.Run()
}
