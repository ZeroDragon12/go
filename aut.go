package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Pass_word"
	dbname   = "postgres"
)

type User struct {
	GUID string   `json:"guid"`
	Rts  []string `json:"rts"`
}

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println(connStr)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println(connStr)
	fmt.Println("Connected!")
	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		return
	}
	User

}
