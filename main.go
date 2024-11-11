package main

//go:generate /home/jakob/go/bin/swag init -g main.go

import (
	"net/http"
	"test/database"
	"test/wanikani"

	_ "test/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Your API Title
// @version 1.0
// @description This is a sample server for managing streaks, subjects, and review statistics.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
	database.ConnectDatabase()

	router.GET("/streak", LongestStreak)

	router.GET("/subjects/:id", Subjects)

	router.GET("/review_statistics/:id", ReviewStatisticsByID)

	router.GET("/review_statistics", ReviewAllStatistics)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}

// LongestStreak handles retrieving the longest streak for the user.
// @Summary Get Longest Streak
// @Description Retrieves the user's longest streak.
// @Tags Streak
// @Produce json
// @Success 200 {object} map[string]string "message: Longest streak data"
// @Router /streak [get]
func LongestStreak(c *gin.Context) {
	body_string := wanikani.GetLongestStreak(c)
	c.JSON(http.StatusOK, gin.H{
		"message": string(body_string),
	})
}

// @Summary Get Subject by ID
// @Description Retrieves information about a specific subject.
// @Tags Subjects
// @Param id path string true "Subject ID"
// @Produce json
// @Success 200 {object} map[string]string "message: Subject information"
// @Router /subjects/{id} [get]
func Subjects(c *gin.Context) {
	id := c.Param("id")
	body_string := wanikani.GetSubjects(c, id)
	c.JSON(http.StatusOK, gin.H{
		"message": string(body_string),
	})
}

// @Summary Get Review Statistics by ID
// @Description Retrieves review statistics for a specific subject.
// @Tags ReviewStatistics
// @Param id path string true "Subject ID"
// @Produce json
// @Success 200 {object} map[string]string "message: Review statistics information"
// @Router /review_statistics/{id} [get]
func ReviewStatisticsByID(c *gin.Context) {
	id := c.Param("id")
	body_string := wanikani.GetReviewStatistics(c, id)
	c.JSON(http.StatusOK, gin.H{
		"message": string(body_string),
	})
}

// @Summary Get All Review Statistics
// @Description Retrieves all review statistics and saves them to the database.
// @Tags ReviewStatistics
// @Produce json
// @Success 200 {object} map[string]string "message: All review statistics information"
// @Router /review_statistics [get]
func ReviewAllStatistics(c *gin.Context) {
	body_string := wanikani.GetReviewStatistics(c, "")
	database.SaveReviewStatisticsToDB(body_string)
	c.JSON(http.StatusOK, gin.H{
		"message": string(body_string),
	})
}
