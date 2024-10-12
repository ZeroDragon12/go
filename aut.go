package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jacks/pgx"
	_ "gorm.io/driver/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "username"
	password = "Pass_word"
	dbname   = "myappdb"
)

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
func insertData(db *sql.DB, name string, age int) error {
	insertSQL := "INSERT INTO users (name, age) VALUES ($1, $2)"

	_, err := db.Exec(insertSQL, name, age)
	return err
}
func main() {
	a := 5
	fmt.Println(a)
	connectDB()
}
