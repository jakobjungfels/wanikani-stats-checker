package wanikani

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// https://mholt.github.io/json-to-go/
type WaniKaniResponse struct {
	Object string `json:"object"`
	URL    string `json:"url"`
	Pages  struct {
		PerPage     float64 `json:"per_page"`
		NextURL     string  `json:"next_url"`
		PreviousURL any     `json:"previous_url"`
	} `json:"pages"`
	TotalCount    float64   `json:"total_count"`
	DataUpdatedAt time.Time `json:"data_updated_at"`
	ReviewEntries []struct {
		ID            float64   `json:"id"`
		Object        string    `json:"object"`
		URL           string    `json:"url"`
		DataUpdatedAt time.Time `json:"data_updated_at"`
		Data          struct {
			CreatedAt            time.Time `json:"created_at"`
			SubjectID            float64   `json:"subject_id"`
			SubjectType          string    `json:"subject_type"`
			MeaningCorrect       float64   `json:"meaning_correct"`
			MeaningIncorrect     float64   `json:"meaning_incorrect"`
			MeaningMaxStreak     float64   `json:"meaning_max_streak"`
			MeaningCurrentStreak float64   `json:"meaning_current_streak"`
			ReadingCorrect       float64   `json:"reading_correct"`
			ReadingIncorrect     float64   `json:"reading_incorrect"`
			ReadingMaxStreak     float64   `json:"reading_max_streak"`
			ReadingCurrentStreak float64   `json:"reading_current_streak"`
			PercentageCorrect    float64   `json:"percentage_correct"`
			Hidden               bool      `json:"hidden"`
		} `json:"data"`
	} `json:"data"`
}

func LongestStreak(context *gin.Context) (body_string_return string) {
	body_string_return = GetReviewStatistics(context, "")

	var highest_streak_meaning float64 = 0
	var highest_streak_meaning_id float64 = 0
	var highest_streak_reading float64 = 0
	var highest_streak_reading_id float64 = 0

	var response WaniKaniResponse
	json.Unmarshal([]byte(body_string_return), &response)

	for _, review_entry := range response.ReviewEntries {
		if highest_streak_meaning < review_entry.Data.MeaningMaxStreak {
			highest_streak_meaning = review_entry.Data.MeaningMaxStreak
			highest_streak_meaning_id = review_entry.Data.SubjectID
			fmt.Println("New highest streak meaning", highest_streak_meaning, "for subject", highest_streak_meaning_id)
		}
		if highest_streak_reading < review_entry.Data.ReadingMaxStreak {
			highest_streak_reading = review_entry.Data.ReadingMaxStreak
			highest_streak_reading_id = review_entry.Data.SubjectID
			fmt.Println("New highest streak reading", highest_streak_reading, "for subject", highest_streak_reading_id)
		}
	}

	body_string_return = strconv.FormatFloat(highest_streak_meaning, 'f', -1, 64)
	body_string_return += "," + strconv.FormatFloat(highest_streak_meaning_id, 'f', -1, 64)
	body_string_return += GetSubjectFields(context, highest_streak_meaning_id)

	body_string_return += ";" + strconv.FormatFloat(highest_streak_reading, 'f', -1, 64)
	body_string_return += "," + strconv.FormatFloat(highest_streak_reading_id, 'f', -1, 64)
	body_string_return += GetSubjectFields(context, highest_streak_reading_id)

	return
}

func GetSubjectFields(context *gin.Context, id float64) (body_string_return string) {
	body_string_subject := GetSubjects(context, strconv.FormatFloat(id, 'f', -1, 64))
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

func GetSubjects(context *gin.Context, id string) (body_string_return string) {
	body_string_return, _ = Get(context, "https://api.wanikani.com/v2/subjects/"+id)
	return
}

func GetReviewStatistics(context *gin.Context, id string) (body_string_return string) {
	body_string_return, _ = Get(context, "https://api.wanikani.com/v2/review_statistics/"+id)
	return
}

func Get(context *gin.Context, url string) (body string, body_byte []byte) {
	var bearer = "Bearer " + os.Getenv("WANIKANI_TOKEN")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	body_byte, err = io.ReadAll(resp.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	resp.Body.Close()
	body = string(body_byte)
	return
}
