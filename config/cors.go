package config

import (
	"github.com/gin-contrib/cors"
)

func initCors() {
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	}

	Cfg.CorsOption = &corsConfig
}
