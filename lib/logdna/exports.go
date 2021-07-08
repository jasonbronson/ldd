package logdna

import (
	"log"
	"net/url"
	"strconv"
)

type Line struct {
	Account  string `json:"_account,omitempty"`
	Cluster  string `json:"_cluster,omitempty"`
	Host     string `json:"_host,omitempty"`
	Ingester string `json:"_ingester,omitempty"`
	Label    struct {
		App               string `json:"app,omitempty"`
		Pod_template_hash string `json:"pod-template-hash,omitempty"`
	} `json:"_label,omitempty"`
	Logtype     string   `json:"_logtype,omitempty"`
	Tag         []string `json:"_tag,omitempty"`
	File        string   `json:"_file,omitempty"`
	Line        string   `json:"_line,omitempty"`
	Rawline     string   `json:"_rawline,omitempty"`
	Level       string   `json:"level,omitempty"`
	TS          int64    `json:"_ts,omitempty"`
	App         string   `json:"_app,omitempty"`
	Pod         string   `json:"pod,omitempty"`
	Namespace   string   `json:"namespace,omitempty"`
	Container   string   `json:"container,omitempty"`
	Containerid string   `json:"containerid,omitempty"`
	Node        string   `json:"node,omitempty"`
	IP          string   `json:"_ip,omitempty"`
	Key         string   `json:"__key,omitempty"`
	Bid         string   `json:"_bid,omitempty"`
	ID          string   `json:"_id,omitempty"`
}

type ExportReponse struct {
	PaginationId string `json:"pagination_id,omitempty"`
	Lines        []Line `json:"lines,omitempty"`
}

func (client *LogdnaClient) GetLog(startTime int64, endTime int64, levels string, query string, apps string) (ExportReponse, error) {
	var exportReponse ExportReponse
	path := "export?" +
		"from=" + strconv.FormatInt(startTime, 10) +
		"&to=" + strconv.FormatInt(endTime, 10) +
		"&levels=" + levels +
		"&query=" + url.QueryEscape(query) +
		"&apps=" + apps

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
