package logdna

import (
	"log"
	"net/url"
	"strconv"
)

type Line struct {
	Account  string `json:"_account"`
	Cluster  string `json:"_cluster"`
	Host     string `json:"_host"`
	Ingester string `json:"_ingester"`
	Label    struct {
		App               string `json:"app"`
		Pod_template_hash string `json:"pod-template-hash"`
	} `json:"_label"`
	Logtype     string   `json:"_logtype"`
	Tag         []string `json:"_tag"`
	File        string   `json:"_file"`
	Line        string   `json:"_line"`
	Rawline     string   `json:"_rawline"`
	Level       string   `json:"level"`
	TS          int64    `json:"_ts"`
	App         string   `json:"_app"`
	Pod         string   `json:"pod"`
	Namespace   string   `json:"namespace"`
	Container   string   `json:"container"`
	Containerid string   `json:"containerid"`
	Node        string   `json:"node"`
	IP          string   `json:"_ip"`
	Key         string   `json:"__key"`
	Bid         string   `json:"_bid"`
	ID          string   `json:"_id"`
}

type ExportReponse struct {
	PaginationId string `json:"pagination_id"`
	Lines        []Line `json:"lines"`
}

func (client *LogdnaClient) GetLog(startTime int64, endTime int64, levels string, query string) (ExportReponse, error) {
	var exportReponse ExportReponse
	path := "export?" +
		"from=" + strconv.FormatInt(startTime, 10) +
		"&to=" + strconv.FormatInt(endTime, 10) +
		"&levels=" + levels +
		"&query=" + url.QueryEscape(query) +
		"&apps=video-api-cron"

	log.Println(path)

	resp, err := client._get(
		path,
		[]byte(""))
	if err != nil {
		log.Printf("Error GetLog %v", err)
		return exportReponse, err
	}
	err = processResponse(resp, &exportReponse)
	return exportReponse, err
}
