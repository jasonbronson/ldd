package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bugsnag/bugsnag-go"
	"github.com/gin-gonic/gin"
)

type Error struct {
	StatusCode int    `json:"code,omitempty"`
	Title      string `json:"title,omitempty"`
	Detail     string `json:"detail,omitempty"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Success struct {
	Data interface{} `json:"data"`
}

func SendError(g *gin.Context, statusCode int, err error) {
	log.Printf("error code %v, details: %v", statusCode, err.Error())

	errs := ErrorResponse{
		Errors: []Error{
			{
				StatusCode: statusCode,
				Title:      http.StatusText(statusCode),
				Detail:     err.Error(),
			},
		},
	}
	g.AbortWithStatusJSON(statusCode, errs)

}

func SendSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	if statusCode == http.StatusNoContent {
		return
	}

	if statusCode == http.StatusInternalServerError {
		bugsnag.Notify(fmt.Errorf("event#exception=\"Internal Server Error\" description=\"error was: %v\" start_time=%v", data, time.Now().Unix()))
	}

	res := Success{
		Data: data,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		bugsnag.Notify(fmt.Errorf("error marshalling payload: %v", err))
	}

}
