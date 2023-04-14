package config

import (
	"database/sql"
	"fmt"
	"log"
   _ "github.com/lib/pq"	
)
const (
	DB_HOST = "localhost"
	DB_PORT = 5432
	DB_USER = "postgres"
	DB_PASSWORD = "1q2w3e"
	DB_NAME = "belajar"
)

var DB *sql.DB

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	fmt.Println("successfully connected")
	DB = db
}
