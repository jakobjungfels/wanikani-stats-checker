package wanikani

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWaniKaniGet(t *testing.T) {
	type args struct {
		c   *gin.Context
		url string
	}
	tests := []struct {
		name     string
		args     args
		wantBody []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBody := WaniKaniGet(tt.args.c, tt.args.url); !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("WaniKaniGet() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}
