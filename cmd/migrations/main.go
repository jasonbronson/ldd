package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)
func main() {
	// Flag command line
	databaseurl := flag.String("database", os.Getenv("DATABASE_URL"), "a string")
	dir := flag.String("source", os.Getenv("MIGRATIONS_PKG_DIR"), "a string")
	flag.Parse()
	if *databaseurl == "" {
		*databaseurl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))
	}
	if *dir == "" {
		*dir = "./migrations"
	}
	abs, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatal("Path not valid", err)
	}
	_, err = os.Stat(abs)
	if os.IsNotExist(err) {
		log.Fatal("path to migrations is invalid")
	}
	m, err := migrate.New("file://"+abs, *databaseurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Running migrations")
	err = m.Up();
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	if err == migrate.ErrNoChange {
		log.Println("No changes")		
	}

	fmt.Println("Migrations complete")
}
