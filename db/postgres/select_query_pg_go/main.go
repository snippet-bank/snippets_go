package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myappuser"
	password = "asdf1234"
	dbname   = "myapp"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, email, username FROM account LIMIT $1", 10)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var email string
		var username string
		err = rows.Scan(&id, &email, &username)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, email, username)
	}
	// error during scan?
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
