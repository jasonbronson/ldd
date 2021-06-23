package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Logs struct {
	Id              string    `gorm:"primary_key" json:"id"`
	Log             string    `json:"line"`
	Last_error      time.Time `json:"last_error"`
	Updated_at      time.Time `json:"updated_at"`
	Error_count     int64     `json:"error_count"`
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
