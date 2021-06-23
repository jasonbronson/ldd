package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Live(g *gin.Context) {
	g.Writer.WriteHeader(http.StatusOK)
	g.Writer.Write([]byte(fmt.Sprintf("%d %s", http.StatusOK, http.StatusText(http.StatusOK))))
}

func GetLogMatches(g *gin.Context) {
	g.JSON(200, "GetLogMatches")
}

func GetMatches(g *gin.Context) {
	g.JSON(200, "GetMatches")
}

func PostMatches(g *gin.Context) {
	g.JSON(200, "PostMatches")
}

func PatchMatches(g *gin.Context) {
	g.JSON(200, "PatchMatches")
}
