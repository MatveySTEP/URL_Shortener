package main

import (
	"fmt"
	"github.com/MatveySTEP/URL_Shortener/handlers"
	"github.com/MatveySTEP/URL_Shortener/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Start web server",
		})
	})
	r.POST("/Create-short-url", func(c *gin.Context) {
		handlers.CreateShortUrl(c)
	})
	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.HandleShortUrlRedirect(c)
	})

	store.InitStore()

	if err := r.Run(":8000"); err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
