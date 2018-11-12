package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/geniuscasio/rest-gotrello/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	fmt.Println("Message")
	if port == "" {
		fmt.Println("Port must be set")
		port = ":8000"
	} else {
		port = ":" + port
	}
	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Println(dir)
	// This will serve files under http://localhost:8000/static/<filename>
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	router.HandleFunc("/", endpoints.Get).Methods("GET")
	router.HandleFunc("/income/{id}", endpoints.Get).Methods("GET")
	router.HandleFunc("/income/", endpoints.Create).Methods("POST")
	router.HandleFunc("/income/", endpoints.Get).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
