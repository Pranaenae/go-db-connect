package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBHOST") + ":3306",
		DBName: os.Getenv("DBNAME"),
	}
	// Get a database handle.
	log.SetPrefix("Server: ")
	log.SetFlags(log.Ldate | log.Ltime)

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	http.HandleFunc("/", handler)
	port := "8080"
	fmt.Printf("server is listening on port %v", port)
	serverErr := http.ListenAndServe(":"+port, nil)
	if serverErr != nil {
		fmt.Println("Error:", serverErr)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request from", r.RemoteAddr, "at", r.URL.Path)
	fmt.Fprint(w, "Hello, World")
}
