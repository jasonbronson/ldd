package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	if Cfg.DatabaseURL == "" {
		Cfg.DatabaseURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_DB_USER"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_PORT"), os.Getenv("POSTGRES_DB_DATABASE"))
	}

	driver := "postgres"

	db, err := sql.Open(driver, Cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	Cfg.GormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	db.SetMaxOpenConns(Cfg.MaxConnections)
	db.SetMaxIdleConns(Cfg.MaxIdleConnections)
	if err = db.Ping(); err != nil {
		log.Fatal("Master Database cannot connect ", err)
	}
	log.Println("Success connecting to master database")

}
