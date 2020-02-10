package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// in a real app these would come from the environment, via os.Getenv
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

	sqlStatement := `
		INSERT INTO account (email, username)
		VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, "nobunaga@gmail.com", "samuraiguy33")
	if err != nil {
		panic(err)
	}
}
