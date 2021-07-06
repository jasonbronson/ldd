package models

import (
	"time"
)

type LogsFound struct {
	LogsID    string    `gorm:"primary_key" json:"logsID"`
	TimeStart time.Time `json:"time_start"`
	TimeEnd   time.Time `json:"time_end"`
}

func (d LogsFound) TableName() string {
	return "logs_found"
}
