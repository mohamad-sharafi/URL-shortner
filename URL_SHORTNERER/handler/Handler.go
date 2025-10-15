package handler

import (
	"URL_shortnerer/shortner"
	"URL_shortnerer/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UrlCreationRequestBody struct {
	Url    string `json:"url" binding:"required"`
	Userid string `json:"user_id" binding:"required"`
}

func CreateShorturl(c *gin.Context) {
	var creationRequest UrlCreationRequestBody
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	shorturl := shortner.GenerateShortUrl(creationRequest.Url, creationRequest.Userid)
	store.SaveUrlMapping(shorturl, creationRequest.Url, creationRequest.Userid)
	c.JSON(http.StatusOK, gin.H{
		"message":  "short url created successfully",
		"shorturl": "127.0.0.1:8080/" + shorturl,
	})
}

func HandleShorturlRedirect(c *gin.Context) {
	shorturl := c.Param("shorturl")
	initialUrl := store.RetrieveInitialUrl(shorturl)
	if initialUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, initialUrl)
}