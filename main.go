package main

import (
	"url-shortener/config"
	"url-shortener/controllers"
	"url-shortener/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.InitDB()

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/:short_code", controllers.GetOriginalURL)

	r.Run(":8080")
}
