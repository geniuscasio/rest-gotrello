package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	InitDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8005"
	}
	port = ":" + port
	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Println("Running server on port " + port)

	router.Use(SecureEndoint)
	router.Use(AccessLogMiddleware)

	// GETs
	router.HandleFunc("/api/v1/login/", Login).Methods("POST")
	router.HandleFunc("/api/v1/income/", Get).Methods("GET")
	router.HandleFunc("/api/v1/income/{id}", Get).Methods("GET")

	// POSTs
	router.HandleFunc("/api/v1/income/", Create).Methods("POST")

	// Static "/" must be last in code
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(port, router))
}
