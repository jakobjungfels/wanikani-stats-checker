package main

import (
	"net/http"
	"test/database"
	"test/wanikani"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()
	router.GET("/streak", func(c *gin.Context) {
		body_string := wanikani.LongestStreak(c)
		c.JSON(http.StatusOK, gin.H{
			"message": string(body_string),
		})
	})
	router.GET("/subjects/:id", func(c *gin.Context) {
		id := c.Param("id")
		body := wanikani.WaniKaniGet(c, "https://api.wanikani.com/v2/subjects/"+id)
		body_string := string(body)
		c.JSON(http.StatusOK, gin.H{
			"message": string(body_string),
		})
	})
	router.GET("/review_statistics/:id", func(c *gin.Context) {
		id := c.Param("id")
		body := wanikani.WaniKaniGet(c, "https://api.wanikani.com/v2/review_statistics/"+id)
		body_string := string(body)
		c.JSON(http.StatusOK, gin.H{
			"message": string(body_string),
		})
	})
	router.GET("/review_statistics", func(c *gin.Context) {
		body_string := wanikani.GetReviewStatistics(c)
		database.AddReviewStatistics(body_string)
		c.JSON(http.StatusOK, gin.H{
			"message": string(body_string),
		})
	})
	router.Run()
}
