package repository

import (
	"fmt"

	"github.com/jasonbronson/ldd/models"
	"gorm.io/gorm"
)

type MatchesRequest struct {
	Matching_string string `form:"matching_string" json:"matching_string"`
	Name            string `form:"name" json:"name"`
	Apps            string `form:"apps" json:"apps"`
	Description     string `form:"description" json:"description"`
}

func CreateMatches(db *gorm.DB, match models.Matches) error {
	return db.Create(&match).Error
}

func PatchMatches(db *gorm.DB, match models.Matches) error {
	return db.Updates(&match).Error
}

func GetAllMatches(db *gorm.DB) ([]*models.Matches, error) {
	var matches []*models.Matches
	result := db.Find(&matches)
	return matches, result.Error
}

func GetMatchesByMatchString(db *gorm.DB, match_string string) (*models.Matches, error) {
	var match *models.Matches
	fmt.Println(match_string)
	result := db.Where("matching_string=?", match_string).Find(&match)
	return match, result.Error
}

func GetMatchesByID(db *gorm.DB, id string) (models.Matches, error) {
	var match models.Matches
	result := db.Where("id=?", id).Find(&match)
	return match, result.Error
}
