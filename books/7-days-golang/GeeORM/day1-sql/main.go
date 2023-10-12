package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "gee.db")
	check(err)

	defer func() {
		_ = db.Close()
	}()

	_, err = db.Exec("DROP TABLE IF EXISTS User;")
	check(err)
	_, err = db.Exec("CREATE TABLE User(Name text);")
	check(err)

	res, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "Alice", "Ai")
	check(err)
	affected, _ := res.RowsAffected()
	log.Printf("%d rows affected\n", affected)

	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	err = row.Scan(&name)
	check(err)
	log.Println("name:", name)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
