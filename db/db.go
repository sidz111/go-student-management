package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	db_url := "root:root@tcp(localhost:3303)/studentdb"
	db, err := sql.Open("mysql", db_url)
	if err != nil {
		log.Fatal("Db Not Found")
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Not Connect")
	}
	fmt.Println("Db Connected Done")
	return db
}
