package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonbronson/ldd/config"
	"github.com/jasonbronson/ldd/controller"
)

func NewRoute(cfg *config.Config) http.Handler {
	router := gin.Default()

	router.GET("/", controller.Live)

	router.GET("/api/logmatches", controller.GetLogMatches)
	router.GET("/api/matches", controller.GetMatches)

	router.POST("/api/matches", controller.PostMatches)

	router.PATCH("/api/matches", controller.PatchMatches)
	return router
}
