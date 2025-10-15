package main

import (
	"URL_shortnerer/handler"
	"URL_shortnerer/store"
	"log"

	"github.com/gin-gonic/gin"
)

const Port string = ":8080"

func main() {
	store.InitstoreService()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to my URL Shortnerer API",
		})
	})
	r.POST("/create", handler.CreateShorturl)
	r.GET("/:shorturl", handler.HandleShorturlRedirect)

	log.Printf("Starting server on port %v", Port)
	if err := r.Run(Port); err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
