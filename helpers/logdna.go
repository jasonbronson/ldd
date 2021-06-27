package helpers

import (
	"log"

	"github.com/jasonbronson/ldd/config"
	logdnasdk "github.com/jasonbronson/ldd/lib/logdna"
)

var client *logdnasdk.LogdnaClient

func init() {
	client = logdnasdk.New(config.Cfg.ServiceKey)
}

func GetLogs(startTime int64, endTime int64, levels string, query string) (logsLine logdnasdk.ExportReponse, err error) {
	logsLine, err = client.GetLog(startTime, endTime, levels, query)

	if err != nil {
		log.Println("error getting logs ", logsLine)
		return
	}
	log.Println(logsLine)
	return
}
