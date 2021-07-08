package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	if Cfg.DatabaseURL == "" {
		Cfg.DatabaseURL = "file:ldd.db?cache=shared"
	}

	var err error
	Cfg.GormDB, err = gorm.Open(sqlite.Open(Cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	db, err := Cfg.GormDB.DB()

	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	//defer db.Close()
	db.SetMaxOpenConns(Cfg.MaxConnections)
	db.SetMaxIdleConns(Cfg.MaxIdleConnections)
	if err = db.Ping(); err != nil {
		log.Fatal("Master Database cannot connect ", err)
	}
	log.Println("Success connecting to master database")

}
