package config

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var Cfg = &Config{}

type Config struct {
	GormDB             *gorm.DB
	Port               string
	DisableSSL         bool
	CorsOption         *cors.Config
	DatabaseURL        string
	MaxConnections     int
	MaxIdleConnections int
	ServiceKey         string
	Levels             string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	initEnv()
	initCors()
	initDB()

}

func initEnv() {
	Cfg.Port = os.Getenv("PORT")
	Cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	Cfg.MaxConnections, _ = strconv.Atoi(os.Getenv("DATABASE_MAX_CONNECTIONS"))
	Cfg.MaxIdleConnections, _ = strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNECTIONS"))
	Cfg.DisableSSL, _ = strconv.ParseBool(os.Getenv("DISABLE_SSL"))
	Cfg.ServiceKey = os.Getenv("SERVICE_KEY")
	Cfg.Levels = os.Getenv("LEVELS")
}
