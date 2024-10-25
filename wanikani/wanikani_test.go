package wanikani

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestWaniKaniGet(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error occurred while loading .env")
	}
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	type args struct {
		c   *gin.Context
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantBody string
	}{
		{
			name: "TestCaseOne",
			args: args{
				c:   context,
				url: "https://api.wanikani.com/v2/review_statistics",
			},
			wantBody: "https://api.wanikani.com/v2/review_statistics",
		},
		{
			name: "TestCaseTwo",
			args: args{
				c:   context,
				url: "https://api.wanikani.com/v2/subjects",
			},
			wantBody: "https://api.wanikani.com/v2/subjects",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var response WaniKaniResponse
			_, body_byte := Get(test.args.c, test.args.url)
			json.Unmarshal([]byte(body_byte), &response)
			if !reflect.DeepEqual(response.URL, test.wantBody) {
				t.Errorf("Get() = %v, want %v", response.URL, test.wantBody)
			}
		})
	}
}
