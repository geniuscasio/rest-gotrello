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

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expiration := time.Now().Add(365 * 24 * time.Hour)
    	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
		http.SetCookie(w, &cookie)
		fmt.Println("set cookie in logging")
		f(w, r)
	}
}

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

	router.HandleFunc("/", logging(endpoints.Get)).Methods("GET")
	router.HandleFunc("/income/{id}", logging(endpoints.Get)).Methods("GET")
	router.HandleFunc("/income/", logging(endpoints.Create)).Methods("POST")
	router.HandleFunc("/income/", logging(endpoints.Get)).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
