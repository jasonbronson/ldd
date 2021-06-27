package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Logs struct {
	Id              string    `gorm:"primary_key" json:"id"`
	Log_line        string    `json:"log_line"`
	Last_error      time.Time `json:"last_error"`
	Updated_at      time.Time `json:"updated_at"`
	Error_count     int64     `json:"error_count"`
	Time_start      int64     `json:"time_start"`
	Time_end        int64     `json:"time_end"`
	Matching_string string    `json:"matching_string"`
}

func (d Logs) TableName() string {
	return "logs"
}
func (d *Logs) BeforeCreate(tx *gorm.DB) (err error) {
	if d.Id == "" {
		d.Id = uuid.Must(uuid.NewV4(), nil).String()
	}
	return nil
}
