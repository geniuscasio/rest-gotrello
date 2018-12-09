package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	ends "github.com/geniuscasio/rest-gotrello/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Env POST must be set!")
		port = ":8005"
	} else {
		port = ":" + port
	}
	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Println("Running server on port " + port)

	// GETs
	router.HandleFunc("/api/v1/login/", ends.Login).Methods("POST")
	router.HandleFunc("/api/v1/income/", ends.SecureEndpoint(ends.Get)).Methods("GET")
	router.HandleFunc("/api/v1/income/{id}", ends.SecureEndpoint(ends.Get)).Methods("GET")

	// POSTs
	router.HandleFunc("/api/v1/income/", ends.SecureEndpoint(ends.Create)).Methods("POST")

	// Static "/" must be last in code
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(port, router))
}
