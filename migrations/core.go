package migrations

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/jasonbronson/ldd/config"
)

var Migrations = make([]*gormigrate.Migration, 0)

func Run() {
	log.Println("Run migrations complete")
	db := config.Cfg.GormDB
	m := gormigrate.New(db, gormigrate.DefaultOptions, Migrations)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
