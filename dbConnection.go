package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func connectToDb() *sql.DB {
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL not found in env file")
	}

	//Connect to Database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to Database")
	return db
}
