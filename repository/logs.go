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

func GetLogFromMatchingString(db *gorm.DB, logs *models.Logs) (models.Logs, error) {
	var logLine models.Logs
	result := db.Find(&logLine).Where("matching_string=?", logs.MatchingString)
	return logLine, result.Error
}

func CreateLogs(db *gorm.DB, log models.Logs) error {
	return db.Create(&log).Error
}

func UpdateLogs(db *gorm.DB, logs models.Logs) error {
	return db.Model(logs).Updates(&logs).Error
}
