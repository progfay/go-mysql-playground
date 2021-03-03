package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	dbUser     = "MYSQL_USER"
	dbPassword = "MYSQL_PASSWORD"
	dbHost     = "MYSQL_HOSTNAME"
	dbPort     = "MYSQL_PORT"
	dbName     = "MYSQL_DATABASE"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	type Row struct {
		ID int `db:"id"`
	}

	rows := make([]Row, 0)
	err = db.Select(&rows, "SELECT id FROM user")
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
