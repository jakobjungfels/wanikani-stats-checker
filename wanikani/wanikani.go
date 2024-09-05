package wanikani

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LongestStreak(c *gin.Context) (body_string string) {
	body_string = GetReviewStatistics(c)

	var highest_streak float64 = 0
	var result map[string]interface{}
	json.Unmarshal([]byte(body_string), &result)

	for key, value := range result {
		switch value.(type) {
		case interface{}:
			if key == "data" {
				for _, value2 := range value.([]interface{}) {
					for key3, value3 := range value2.(map[string]interface{}) {
						if key3 == "data" {
							for key4, value4 := range value3.(map[string]interface{}) {
								if key4 == "meaning_max_streak" {
									if highest_streak < value4.(float64) {
										highest_streak = value4.(float64)
										fmt.Println("New highest streak", highest_streak)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	body_string = strconv.FormatFloat(highest_streak, 'f', -1, 64)

	return
}

func GetReviewStatistics(c *gin.Context) (body_string string) {
	body := WaniKaniGet(c, "https://api.wanikani.com/v2/review_statistics/")
	body_string = string(body)
	return
}

func WaniKaniGet(c *gin.Context, url string) (body []byte) {
	var bearer = "Bearer " + os.Getenv("WANIKANI_TOKEN")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	resp.Body.Close()
	return
}
