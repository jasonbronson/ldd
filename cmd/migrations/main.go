package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	databaseurl := os.Getenv("DATABASE_URL")
	if databaseurl == "" {
		databaseurl = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_DB_USER"), os.Getenv("POSTGRES_DB_PASSWORD"), os.Getenv("POSTGRES_DB_HOST"), os.Getenv("POSTGRES_DB_PORT"), os.Getenv("POSTGRES_DB_DATABASE"))
	}

	dir := "./migrations"
	abs, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal("Path not valid", err)
	}
	_, err = os.Stat(abs)
	if os.IsNotExist(err) {
		log.Fatal("path to migrations is invalid")
	}
	m, err := migrate.New("file://"+abs, databaseurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Running migrations")
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	if err == migrate.ErrNoChange {
		log.Println("No changes")
	}

	fmt.Println("Migrations complete")
}
