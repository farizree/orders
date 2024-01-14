package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Environment() (string, error) {
	env := "development"

	return env, nil
}
func Hostname() (string, error) {
	host := "localhost"

	return host, nil
}

func DetermineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2010"
	}
	return ":" + port, nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Bussan100"
	dbname   = "payme"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	checkErr(err)

	return db

	// fmt.Println("Successfully connected!")
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err.Error())
	}
}
