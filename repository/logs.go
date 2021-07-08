package repository

import (
	"github.com/jasonbronson/ldd/models"
	"gorm.io/gorm"
)

func GetLogsLine(db *gorm.DB) ([]*models.Logs, error) {

	var logsLine []*models.Logs
	result := db.Find(&logsLine)
	return logsLine, result.Error
}

func GetLogsByMatchingString(db *gorm.DB, matchingString string) (*models.Logs, error) {
	var logLine *models.Logs
	result := db.Where("matching_string=?", matchingString).Find(&logLine)
	return logLine, result.Error
}

func CreateLogs(db *gorm.DB, log models.Logs) error {
	return db.Create(&log).Error
}

func UpdateLogs(db *gorm.DB, logs models.Logs) error {
	return db.Model(logs).Updates(&logs).Error
}
