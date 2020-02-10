package main

import (
	"database/sql"
	"fmt"

	// The underscore means we are importing this package to register postgres drivers
	// but will not reference this package directly.
	_ "github.com/lib/pq"
)

// Assumes you have followed instructions in "postgres-setup" directory.
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

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
