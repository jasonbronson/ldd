package repository

import (
	"github.com/jasonbronson/ldd/models"
	"gorm.io/gorm"
)

func GetLogsFoundById(db *gorm.DB, Id string) (*models.LogsFound, error) {
	var logsFound *models.LogsFound
	result := db.Where("logs_id = ?", Id).First(&logsFound)
	return logsFound, result.Error
}

func CreateLogsFound(db *gorm.DB, logs models.LogsFound) error {
	return db.Create(&logs).Error
}

func UpdateLogsFound(db *gorm.DB, logs models.LogsFound) error {
	return db.Model(logs).Updates(&logs).Error
}
