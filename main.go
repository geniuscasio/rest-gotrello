package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	ends "github.com/geniuscasio/rest-gotrello/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func getDB() *sql.DB {
	if db == nil {
		fmt.Printf("connect to db")
		dbinfo := os.Getenv("DB_INFO")
		dbinfo = "postgres://nhshglygkfiair:5d42deb354a442697dea6593f1f2bb6e0e869bda02adaad22d3cdcbd321671e1@ec2-23-21-122-141.compute-1.amazonaws.com:5432/d2dorvldn0g3nl"
		var err error
		db, err = sql.Open("postgres", dbinfo)
		if err != nil {
			fmt.Printf(err.Error())
			return nil
		}
	}
	return db
}

func initDB() {
	c := getDB()
	_, r := c.Exec(`
	CREATE TABLE MY_USERS(
		user_id serial PRIMARY KEY, 
		username VARCHAR (100) UNIQUE NOT NULL, 
		password VARCHAR (100) NOT NULL
	)`)
	if r != nil {
		fmt.Println(r.Error())
	}
}

func main() {
	router := mux.NewRouter()
	initDB()
	port := ":8005"

	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Println("Running server on port " + port)

	router.Use(ends.SecureEndoint)
	router.Use(ends.AccessLogMiddleware)

	// GETs
	router.HandleFunc("/api/v1/login/", ends.Login).Methods("POST")
	router.HandleFunc("/api/v1/income/", ends.Get).Methods("GET")
	router.HandleFunc("/api/v1/income/{id}", ends.Get).Methods("GET")

	// POSTs
	router.HandleFunc("/api/v1/income/", ends.Create).Methods("POST")

	// Static "/" must be last in code
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(port, router))
}
