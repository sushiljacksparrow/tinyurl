package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tinyurl/models"
	"log"
	"net/http"
	"regexp"
	"time"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()

	router.POST("/url/tiny", func(c *gin.Context) {
		var input models.CreateTinyURLInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tinyUrl := generateAndStoreTinyUrl(input.OriginalURL, 0, input.User)


		log.Println("Created short url for ", tinyUrl)
		c.JSON(http.StatusOK, gin.H{"tiny_url": tinyUrl})
	})

	router.GET("/url/long", func(c *gin.Context) {
		var input models.GetLongURLInput
		if err := c.ShouldBindQuery(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		log.Println("Received query ", input.TinyUrl, input.User)


		var url models.URL
		if err := models.DB.First(&url, "tiny_url = ?", input.TinyUrl).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "Record not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"url": url})
	})


	router.Run()
}

func generateAndStoreTinyUrl(originalURL string, startIndex int, user string) string {
	byteURLData := []byte(originalURL)
	hashedURLData := fmt.Sprintf("%x", md5.Sum(byteURLData))
	tinyURLRegex, err := regexp.Compile("[/+]")
	if err != nil {
		return "Unable to generate tiny URL"
	}
	tinyURLData := tinyURLRegex.ReplaceAllString(base64.URLEncoding.EncodeToString([]byte(hashedURLData)), "_")
	log.Println(tinyURLData)
	if len(tinyURLData) < (startIndex + 6) {
		return "Unable to generate tiny URL"
	}

	tinyURL := tinyURLData[startIndex : startIndex+6]
	var dbURLData models.URL
	models.DB.Where("tiny_url = ?", tinyURL).Find(&dbURLData)
	if dbURLData.TinyURL == "" {
		models.DB.Create(models.URL{OriginalURL: originalURL, TinyURL: tinyURL, CreatedAt: time.Now(), User: user, ID: time.Now().Unix()})
		return tinyURL
	} else if dbURLData.TinyURL == tinyURL && dbURLData.OriginalURL == originalURL {
		log.Println("Url already exists ", tinyURLData)
		return tinyURL
	} else {
		return generateAndStoreTinyUrl(originalURL, startIndex + 1, user)
	}
}
