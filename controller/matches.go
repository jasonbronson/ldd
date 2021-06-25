package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/ldd/config"
	"github.com/jasonbronson/ldd/helpers"
	"github.com/jasonbronson/ldd/models"
	"github.com/jasonbronson/ldd/repository"
)

func Live(g *gin.Context) {
	g.Writer.WriteHeader(http.StatusOK)
	g.Writer.Write([]byte(fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK))))
}

func GetAllMatches(g *gin.Context) {
	db := config.Cfg.GormDB
	db = db.WithContext(g)

	resp, err := repository.GetAllMatches(db)

	if err != nil {
		helpers.SendError(g, http.StatusBadRequest, err)
		return
	}

	helpers.SendSuccess(g.Writer, http.StatusOK, resp)
}

func PostMatches(g *gin.Context) {
	db := config.Cfg.GormDB
	db = db.WithContext(g)

	var requestData *repository.MatchesRequest
	var match models.Matches

	if err := g.ShouldBindJSON(&requestData); err != nil {
		helpers.SendError(g, http.StatusInternalServerError, fmt.Errorf("unable to marshal JSON: %v", err))
		return
	}

	fmt.Println(requestData.Matching_string)

	match.Matching_string = requestData.Matching_string
	match.Name = requestData.Name
	match.Description = requestData.Description

	if match.Name == "" || match.Description == "" {
		helpers.SendError(g, http.StatusBadRequest, errors.New("name or description was empty"))
		return
	}

	if match.Matching_string == "" {
		helpers.SendError(g, http.StatusBadRequest, errors.New("matching_string was empty"))
		return
	}

	fmt.Println(match.Matching_string)
	data, _ := repository.GetMatchFromMatchString(db, match.Matching_string)

	if data.Matching_string != "" {
		helpers.SendError(g, http.StatusBadRequest, errors.New("matching_string already exists"))
		return
	}

	err := repository.CreateMatch(db, match)
	if err != nil {
		helpers.SendError(g, http.StatusBadRequest, err)
		return
	}
	helpers.SendSuccess(g.Writer, http.StatusOK, "SUCCESS")

}

func PatchMatches(g *gin.Context) {
	db := config.Cfg.GormDB
	db = db.WithContext(g)

	var requestData *repository.MatchesRequest
	var match models.Matches

	matchId := g.Param("matchID")

	if err := g.ShouldBindJSON(&requestData); err != nil {
		helpers.SendError(g, http.StatusInternalServerError, fmt.Errorf("unable to marshal JSON: %v", err))
		return
	}

	match.Matching_string = requestData.Matching_string
	match.Name = requestData.Name
	match.Description = requestData.Description
	match.Id = matchId

	if matchId == "" {
		helpers.SendError(g, http.StatusNotFound, fmt.Errorf("matchid was empty"))
		return
	}

	if match.Name == "" || match.Description == "" {
		helpers.SendError(g, http.StatusBadRequest, errors.New("name or description was empty"))
		return
	}

	if match.Matching_string == "" {
		helpers.SendError(g, http.StatusBadRequest, errors.New("matching_string was empty"))
		return
	}

	data, _ := repository.GetMatchFromID(db, match.Id)

	if data.Id == "" {
		helpers.SendError(g, http.StatusBadRequest, errors.New("matchId does not exist"))
		return
	}

	err := repository.PatchMatch(db, match)
	if err != nil {
		helpers.SendError(g, http.StatusBadRequest, err)
		return
	}
	helpers.SendSuccess(g.Writer, http.StatusOK, "SUCCESS")
}
