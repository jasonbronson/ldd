package jobs

import (
	"encoding/json"
	"log"
	"regexp"
	"time"

	"github.com/jasonbronson/ldd/config"
	"github.com/jasonbronson/ldd/helpers"
	"github.com/jasonbronson/ldd/models"
	"github.com/jasonbronson/ldd/repository"
)

const (
	GetLogJobInterval = "@every 0h1m0s"
)

func GetLogJob() {

	endTime := time.Now().Unix()
	startTime := endTime - 432000
	levels := config.Cfg.Levels

	allMatches, err := repository.GetAllMatches(config.Cfg.GormDB)
	if err != nil {
		log.Println(err)
		return
	}

	if len(allMatches) < 1 {
		log.Println("There are no MatchingString in the Matches table")
		return
	}

	for _, item := range allMatches {

		match := item.MatchingString
		log.Printf("Loading logs from logdna using query: %v \n", match)
		data, err := helpers.GetLogs(startTime, endTime, levels, match)
		if err != nil {
			log.Println(err)
			return
		}
		// fmt.Println(data)
		for _, i := range data.Lines {
			if CheckMatchStringAgainstLine(i.Line, match) {
				logLine := &models.Logs{
					LogLine:        i.Line,
					UpdatedAt:      time.Now(),
					LastError:      time.Unix(i.TS/1000, 0),
					MatchingString: match,
				}
				logFound := &models.LogsFound{
					LogsID:    i.ID,
					TimeStart: time.Unix(startTime, 0),
					TimeEnd:   time.Unix(endTime, 0),
				}

				UpdateDB(logLine, logFound)
			}

		}

	}

	log.Println("=====Cronjob Done=====")
}

func UpdateDB(logs *models.Logs, logsFound *models.LogsFound) {

	logsDB, _ := repository.GetLogsByMatchingString(config.Cfg.GormDB, logs.MatchingString)
	if len(logsDB.ID) < 1 {
		//Print the logs
		logsJson, _ := json.Marshal(logs)
		log.Println(string(logsJson))

		err := repository.CreateLogs(config.Cfg.GormDB, *logs)
		if err != nil {
			log.Println("Error to CreateLogs: ", err)
			//repository.UpdateLogs(config.Cfg.GormDB, *logs)
		}
	}

	logsFoundDB, _ := repository.GetLogsFoundById(config.Cfg.GormDB, logsDB.ID)
	if len(logsFoundDB.LogsID) < 1 {
		//Print the logs_found
		logsFoundJson, _ := json.Marshal(logsFound)
		log.Println(string(logsFoundJson))

		logsFound.LogsID = logsDB.ID
		err := repository.CreateLogsFound(config.Cfg.GormDB, *logsFound)
		if err != nil {
			log.Println("Error to CreateLogsFound: ", err)
			// repository.UpdateLogsFound(config.Cfg.GormDB, *logsFound)
		}
	}

}

func CheckMatchStringAgainstLine(line, matchString string) bool {
	regex := regexp.MustCompile(matchString + `\b`)
	return regex.MatchString(line)
}
