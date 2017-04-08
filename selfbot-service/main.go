package main

import (
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	// dsn, err := ioutil.ReadFile("dsn.txt")
	// db, err = sql.Open("mysql", string(dsn))

	// if err != nil {
	// 	panic(err.Error())
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err.Error())
	// }
}

func main() {
	defer db.Close()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":1234", router))
}
