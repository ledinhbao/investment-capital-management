package main

import (
	"./investor"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// TODO Install gcc and add to %PATH% in order to run sqlite3 driver.

func main() {

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlSelect := `SELECT * FROM investors`
	query, err := db.Prepare(sqlSelect)
	if err != nil {
		panic(err)
	}
	var a investor.Investor
	err = query.QueryRow("1").Scan(&a)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)

	fmt.Println("Hello world!")
}
