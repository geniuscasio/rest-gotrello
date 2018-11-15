package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/geniuscasio/rest-gotrello/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, _ := r.Cookie("username")
		fmt.Println(username)

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username", Value: time.Now().Format("2006-01-02 15:04:05"), Expires: expiration}
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

	// GETs
	router.HandleFunc("/api/v1/", logging(endpoints.Get)).Methods("GET")
	router.HandleFunc("/api/v1/income/", logging(endpoints.Get)).Methods("GET")
	router.HandleFunc("/api/v1/income/{id}", logging(endpoints.Get)).Methods("GET")

	// POSTs
	router.HandleFunc("/api/v1/income/", logging(endpoints.Create)).Methods("POST")

	// Static "/" must be last in code
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(port, router))
}
