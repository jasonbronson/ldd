package repository

import (
	"github.com/jasonbronson/ldd/models"
	"gorm.io/gorm"
)

func GetLogsLine(db *gorm.DB, limit int) ([]*models.Logs, error) {

	var logsLine []*models.Logs
	var result *gorm.DB
	if limit != 0 {
		result = db.Order("error_count desc").Limit(limit).Find(&logsLine)
	} else {
		result = db.Find(&logsLine)
	}
	return logsLine, result.Error
}

func GetLogFromMatchingString(db *gorm.DB, matching_string string) (models.Logs, error) {

	var logLine models.Logs

	result := db.Find(&logLine).Where("matching_string=?", matching_string)
	return logLine, result.Error
}

func CreateLogs(db *gorm.DB, log models.Logs) error {
	return db.Create(&log).Error
}

func UpdateLogs(db *gorm.DB, log models.Logs) error {
	return db.Table("logs").Where("matching_string=?", log.Matching_string).Update("error_count", log.Error_count+1).Error
}
