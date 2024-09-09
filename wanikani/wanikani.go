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

func LongestStreak(c *gin.Context) (body_string_return string) {
	body_string_return = GetReviewStatistics(c, "")

	var highest_streak_meaning float64 = 0
	var highest_streak_meaning_id float64 = 0
	var highest_streak_reading float64 = 0
	var highest_streak_reading_id float64 = 0
	var result map[string]interface{}
	json.Unmarshal([]byte(body_string_return), &result)

	for key, value := range result {
		switch value.(type) {
		case interface{}:
			if key == "data" {
				for _, value2 := range value.([]interface{}) {
					for key3, value3 := range value2.(map[string]interface{}) {
						if key3 == "data" {
							for key4, value4 := range value3.(map[string]interface{}) {
								if key4 == "meaning_max_streak" {
									if highest_streak_meaning < value4.(float64) {
										highest_streak_meaning = value4.(float64)
										highest_streak_meaning_id = value3.(map[string]interface{})["subject_id"].(float64)
										fmt.Println("New highest streak meaning", highest_streak_meaning, "for subject", highest_streak_meaning_id)
									}
								} else if key4 == "reading_max_streak" {
									if highest_streak_reading < value4.(float64) {
										highest_streak_reading = value4.(float64)
										highest_streak_reading_id = value3.(map[string]interface{})["subject_id"].(float64)
										fmt.Println("New highest streak reading", highest_streak_reading, "for subject", highest_streak_reading_id)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	body_string_return = strconv.FormatFloat(highest_streak_meaning, 'f', -1, 64)
	body_string_return += "," + strconv.FormatFloat(highest_streak_meaning_id, 'f', -1, 64)
	body_string_return += GetSubjectFields(c, highest_streak_meaning_id)

	body_string_return += ";" + strconv.FormatFloat(highest_streak_reading, 'f', -1, 64)
	body_string_return += "," + strconv.FormatFloat(highest_streak_reading_id, 'f', -1, 64)
	body_string_return += GetSubjectFields(c, highest_streak_reading_id)

	return
}

func GetSubjectFields(c *gin.Context, id float64) (body_string_return string) {
	body_string_subject := GetSubjects(c, strconv.FormatFloat(id, 'f', -1, 64))
	var result_subject map[string]interface{}
	json.Unmarshal([]byte(body_string_subject), &result_subject)
	for key, value := range result_subject {
		switch value.(type) {
		case interface{}:
			if key == "data" {
				for key2, value2 := range value.(map[string]interface{}) {
					if key2 == "meanings" || key2 == "readings" {
						for _, value3 := range value2.([]interface{}) {
							for key4, value4 := range value3.(map[string]interface{}) {
								if key4 == "meaning" || key4 == "reading" {
									body_string_return += "," + value4.(string)
								}
							}
						}
					} else if key2 == "characters" {
						body_string_return += "," + value2.(string)
					}
				}
			}
		}
	}
	return
}

func GetSubjects(c *gin.Context, id string) (body_string_return string) {
	body := Get(c, "https://api.wanikani.com/v2/subjects/"+id)
	body_string_return = string(body)
	return
}

func GetReviewStatistics(c *gin.Context, id string) (body_string_return string) {
	body := Get(c, "https://api.wanikani.com/v2/review_statistics/"+id)
	body_string_return = string(body)
	return
}

func Get(c *gin.Context, url string) (body []byte) {
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
