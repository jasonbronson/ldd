package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/ldd/config"
	"github.com/jasonbronson/ldd/helpers"
	"github.com/jasonbronson/ldd/repository"
)

func GetLogMatches(g *gin.Context) {
	db := config.Cfg.GormDB
	db = db.WithContext(g)

	resp, err := repository.GetLogsLine(db)

	if err != nil {
		helpers.SendError(g.Writer, http.StatusBadRequest, err)
		return
	}

	helpers.SendSuccess(g.Writer, http.StatusOK, resp)
}
