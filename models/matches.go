package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Matches struct {
	ID             string `gorm:"primary_key" json:"id"`
	MatchingString string `json:"matching_string"`
	Apps           string `json:"apps"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}

func (d Matches) TableName() string {
	return "matches"
}
func (d *Matches) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		d.ID = uuid.Must(uuid.NewV4(), nil).String()
	}
	return nil
}
