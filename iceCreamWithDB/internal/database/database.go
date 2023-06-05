package database

import (
	"fmt"

	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	var err error

	const (
		DB_HOST        = "localhost"
		DB_PORT        = 5433
		DB_NAME        = "IceCream"
		DB_USER        = "postgres"
		DB_PASSWORD    = "12345"
		DB_SSLMODE     = "disable"
		DB_CONNTIMEOUT = 10
	)

	dbinfo := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d",
		DB_HOST, DB_PORT, DB_NAME, DB_USER, DB_PASSWORD, DB_SSLMODE, DB_CONNTIMEOUT)

	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println("Database connection successful...")

	return db
}
