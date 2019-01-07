package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

var db *sql.DB

func handler(w http.ResponseWriter, r *http.Request) {
	user := GetUser()
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func GetUser() User {
	user := User{}
	if err := db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&user.Name); err != nil {
		log.Fatal(err)
	}

	return user
}
