package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       string `json:"data,omitempty"`
}

type Success struct {
	Data interface{} `json:"data"`
}

func SendError(w http.ResponseWriter, statusCode int, err error) {
	errorMessage := Error{
		StatusCode: statusCode,
		Message:    http.StatusText(statusCode),
		Data:       err.Error(),
	}

	response, err := json.Marshal(errorMessage)
	if err != nil {
		log.Println("Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}

func SendSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	payload := Success{
		Data: data,
	}

	response, err := json.Marshal(payload)

	if err != nil {
		log.Println("Response Error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}
