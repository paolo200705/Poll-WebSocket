package main

import (
	"log"

	"github.com/Otavio-Fina/live-websocket/middleware"
	"github.com/Otavio-Fina/live-websocket/routes"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/auth/login", middleware.LoginHandler)

	router.POST("/poll", middleware.AuthMidleware(), routes.PostPoll)
	router.GET("/poll", middleware.AuthMidleware(), routes.GetAllPolls)
	router.GET("/poll/:pollID", middleware.AuthMidleware(), routes.GetPoll)
	router.GET("/ws/poll/:pollID", middleware.AuthMidleware(), routes.ConectPoll)

	router.Run(":8080")
}
