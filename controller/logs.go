package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/ldd/config"
	"github.com/jasonbronson/ldd/helpers"
	"github.com/jasonbronson/ldd/repository"
)

func GetLogMatches(g *gin.Context) {
	db := config.Cfg.GormDB
	db = db.WithContext(g)

	defaultLimit := 20
	var limitInt int

	limit := g.Request.URL.Query().Get("limit")
	limitInt, _ = strconv.Atoi(limit)

	if limit == "" {
		limitInt = defaultLimit
	}

	resp, err := repository.GetLogsLine(db, limitInt)

	if err != nil {
		helpers.SendError(g, http.StatusBadRequest, err)
		return
	}

	helpers.SendSuccess(g.Writer, http.StatusOK, resp)
}
