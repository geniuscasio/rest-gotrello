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
		sessionID := "empty"
		userName := "empty"
		cookieSession, _ := r.Cookie("sessionID")
		if cookieSession != nil {
			sessionID = cookieSession.Value
		}

		cookieUser, _ := r.Cookie("userName")
		if cookieUser != nil {
			userName = cookieUser.Value
		}

		fmt.Printf("User cookie userName=%s sessionID=%s \n", userName, sessionID)
		message, ok := endpoints.CheckSession(userName, sessionID)
		fmt.Printf("CheckSession return '%t' with message '%s' \n", ok, message)
		if ok {
			f(w, r)
		} else {
			fmt.Println("ok false")
			http.Redirect(w, r, "/", 301)
		}
	}
}

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Env POST must be set!")
		port = ":8000"
	} else {
		port = ":" + port
	}
	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Println("Running server on port " + port)
	// GETs
	router.HandleFunc("/api/v1/login/", endpoints.Login).Methods("POST")
	router.HandleFunc("/api/v1/", logging(endpoints.Get)).Methods("GET")
	router.HandleFunc("/api/v1/income/", logging(endpoints.Get)).Methods("GET")
	router.HandleFunc("/api/v1/income/{id}", logging(endpoints.Get)).Methods("GET")

	// POSTs
	router.HandleFunc("/api/v1/income/", logging(endpoints.Create)).Methods("POST")

	// Static "/" must be last in code
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(port, router))
}
