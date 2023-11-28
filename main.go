package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	client := 208

	// напишите код здесь
}
