package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var migrationsPath string
	flag.StringVar(&migrationsPath, "migrations", "./internal/adapter/storage/sqlite/migrations/", "path to migrations")
	flag.Parse()

	log.Println("running sqlite migrations")
	db, err := sql.Open("sqlite3", "./db/main.db")
	if err != nil {
		log.Fatalln("error opening sqlite db:", err)
	}
	defer db.Close()

	f, err := os.Open(migrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()

	for _, v := range files {
		if v.IsDir() {
			continue
		}

		contents, err := os.ReadFile(migrationsPath + v.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = tx.Exec(string(contents))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("sqlite migrations finished")
}
