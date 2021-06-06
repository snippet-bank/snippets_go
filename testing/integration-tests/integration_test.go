//+build integration

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

type row struct {
	Id    int32
	Email string
	Name  string
}

var testData = []*row{
	{1, "user@domain.com", "John Doe"},
	{2, "user2@domain.com", "Anna Smith"},
}

var sqlDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqlDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panicf("DB connection failed %v", err)
	}

	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		log.Panicf("DB connection failed %v", err)
	}

	os.Exit(m.Run())
}

func TestDb(t *testing.T) {
	setTestData()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows, err := getRows(db)
	assert.NoError(t, err)
	assert.ElementsMatch(t, rows, testData)

	clearDB()

}

func setTestData() {
	for _, testRow := range testData {
		res, err := sqlDB.Exec(`
INSERT INTO users (id, email, name)
VALUES ($1, $2, $3);
`, testRow.Id, testRow.Email, testRow.Name)
		if err != nil {
			panic(err)
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			log.Fatalf("setTestData - rows affected error: %s", err)
		}

		if rowsAffected != 1 {
			log.Fatalf("setTestData - expected rows affected 1; got %d", rowsAffected)
		}
	}
}

func clearDB() {
	_, err := sqlDB.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}

func getRows(db *sql.DB) ([]*row, error) {
	rows, err := db.Query(`
SELECT id, email, name
FROM users
`)
	if err != nil {
		return nil, err
	}

	var results []*row
	for rows.Next() {
		result := row{}

		err = rows.Scan(&result.Id, &result.Email, &result.Name)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}

	return results, nil
}
