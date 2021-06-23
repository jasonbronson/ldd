package logdna

import (
	"net/http"
)

type LogdnaClient struct {
	Client     *http.Client
	ServiceKey string
}

func New(serviceKey string) *LogdnaClient {
	return &LogdnaClient{Client: &http.Client{}, ServiceKey: serviceKey}
}
