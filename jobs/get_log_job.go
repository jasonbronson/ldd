package jobs

import (
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

	for _, item := range allMatches {

		match := item.MatchingString
		log.Printf("Loading logs from logdna using query: %v \n", match)
		data, err := helpers.GetLogs(startTime, endTime, levels, match)
		if err != nil {
			log.Println(err)
			return
		}
		//fmt.Println(data)
		for _, i := range data.Lines {
			if CheckMatchStringAgainstLine(i.Line, match) {
				log.Println(i.Line)
				logLine := &models.Logs{
					Log_line:       i.Line,
					Updated_at:     time.Now(),
					Last_error:     time.Unix(i.Ts/1000, 0),
					MatchingString: match,
				}
				logFound := &models.LogsFound{
					LogsID:    "",
					TimeStart: time.Unix(startTime, 0),
					TimeEnd:   time.Unix(endTime, 0),
				}

				log.Println("=====Update Database=====")
				UpdateDB(logLine, logFound)
			}

		}

	}

	log.Println("=====Cronjob Done=====")
}

func UpdateDB(logs *models.Logs, logsFound *models.LogsFound) {

	logsDB, _ := repository.GetLogFromMatchingString(config.Cfg.GormDB, logs)
	if len(logsDB.Id) < 1 {
		err := repository.CreateLogs(config.Cfg.GormDB, *logs)
		if err != nil {
			//repository.UpdateLogs(config.Cfg.GormDB, *logs)
		}
	}

	logsFound, _ = repository.GetLogsFoundById(config.Cfg.GormDB, logsDB.Id)
	if len(logsFound.LogsID) < 1 {
		logsFound.LogsID = logsDB.Id
		repository.CreateLogsFound(config.Cfg.GormDB, *logsFound)
		// if err != nil {
		// 	repository.UpdateLogsFound(config.Cfg.GormDB, *logsFound)
		// }
	}

}

func CheckMatchStringAgainstLine(line, matchString string) bool {
	regex := regexp.MustCompile(matchString + `\b`)
	return regex.MatchString(line)
}
