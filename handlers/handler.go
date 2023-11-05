package handlers

import (
	"github.com/MatveySTEP/URL_Shortener/store"
	"github.com/gin-gonic/gin"
)

func CreateShortUrl(c *gin.Context) {
	//заглушка
}
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
