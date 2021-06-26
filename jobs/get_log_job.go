package jobs

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/jasonbronson/ldd/config"
	"github.com/jasonbronson/ldd/helpers"
	"github.com/jasonbronson/ldd/models"
	"github.com/jasonbronson/ldd/repository"
)

const (
	GetLogJobInterval = "@every 1h0m0s"
)

func GetLogJob() {

	startTime := time.Now().Unix()
	endTime := startTime - 3600

	levels := config.Cfg.Levels

	allMatches, err := repository.GetAllMatches(config.Cfg.GormDB)
	if err != nil {
		log.Println(err)
		return
	}

	match_strings := helpers.ConvertMatchestoMatchString(allMatches)

	query := strings.Join(match_strings, " OR ")

	fmt.Println("=====Loading logs from logdna=====")
	data, err := helpers.GetLogs(startTime, endTime, levels, query)

	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(data)

	var logs []*models.Logs
	for _, i := range data.Lines {
		matchString := GetMatchString(i.Line, match_strings)
		log := &models.Logs{
			Error_count:     1,
			Log_line:        i.Line,
			Updated_at:      time.Now(),
			Last_error:      time.Unix(i.Ts/1000, 0),
			Matching_string: matchString,
		}
		logs = append(logs, log)
	}
	fmt.Println("=====Update Database=====")
	UpdateDataBase(logs)
	fmt.Println("=====Cronjob Done=====")
}

func UpdateDataBase(logs []*models.Logs) {
	logMatches, err := repository.GetLogsLine(config.Cfg.GormDB, 0)
	if err != nil {
		log.Println(err)
		return
	}

	var temp string
	match_strings := helpers.ConvertLogMatchestoMatchString(logMatches)
	for _, log := range logs {
		if !Checkstrings(match_strings, log.Matching_string) {
			repository.CreateLogs(config.Cfg.GormDB, *log)
			match_strings = append(match_strings, log.Matching_string)
		}

		// Update Error_count
		if strings.Compare(temp, log.Matching_string) == 0 {
			continue
		}
		data, _ := repository.GetLogFromMatchingString(config.Cfg.GormDB, log.Matching_string)
		repository.UpdateLogs(config.Cfg.GormDB, data)
		temp = log.Matching_string
	}

}

func Checkstrings(slice []string, str string) bool {
	for i := range slice {
		if slice[i] == str {
			return true
		}
	}
	return false
}

func GetMatchString(line string, match_strings []string) string {
	for _, i := range match_strings {
		regex := regexp.MustCompile(i + `\b`)
		if regex.MatchString(line) {
			return i
		}
	}
	return ""
}
