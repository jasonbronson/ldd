package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jasonbronson/ldd/migrations"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Cannot find .env file skipping ")
	}

	dir := ""
	if len(os.Args) != 2 {
		fmt.Println("Migration directory not found, attempting default")
		dir = "./migrations"
	} else {
		dir = os.Args[1]
	}
	abs, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal("Path not valid", err)
	}
	_, err = os.Stat(abs)
	if os.IsNotExist(err) {
		log.Fatal("path to migrations is invalid ", dir)
	}

	migrations.Run()
	fmt.Println("Migrations complete")

}
