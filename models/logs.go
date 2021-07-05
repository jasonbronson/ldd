package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Logs struct {
	ID             string    `gorm:"primary_key" json:"id"`
	LogLine        string    `json:"log_line"`
	LastError      time.Time `json:"last_error"`
	UpdatedAt      time.Time `json:"updated_at"`
	MatchingString string    `json:"matching_string"`
}

func (d Logs) TableName() string {
	return "logs"
}
func (d *Logs) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		d.ID = uuid.Must(uuid.NewV4(), nil).String()
	}
	return nil
}
