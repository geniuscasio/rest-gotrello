package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	dbHelper "github.com/geniuscasio/rest-gotrello/db"
	ends "github.com/geniuscasio/rest-gotrello/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	dbHelper.InitDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8005"
	}
	port = ":" + port
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
