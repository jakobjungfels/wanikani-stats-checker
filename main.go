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
		body_string := wanikani.GetSubjects(c, id)
		c.JSON(http.StatusOK, gin.H{
			"message": string(body_string),
		})
	})
	router.GET("/review_statistics/:id", func(c *gin.Context) {
		id := c.Param("id")
		body_string := wanikani.GetReviewStatistics(c, id)
		c.JSON(http.StatusOK, gin.H{
			"message": string(body_string),
		})
	})
	router.GET("/review_statistics", func(c *gin.Context) {
		body_string := wanikani.GetReviewStatistics(c, "")
		database.SaveReviewStatisticsToDB(body_string)
		c.JSON(http.StatusOK, gin.H{
			"message": string(body_string),
		})
	})
	router.Run()
}
