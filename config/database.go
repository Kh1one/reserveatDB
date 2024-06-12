package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/reserveat")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("db connected")
		DB = db
	}
}
