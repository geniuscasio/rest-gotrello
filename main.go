package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	ends "github.com/geniuscasio/rest-gotrello/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type inMsg struct {
	Phone string `json:"phone"`
}

type outMsg struct {
	Phone  string `json:"phone"`
	Status string `json:"status"`
}

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
	
	router.HandleFunc("/test", sayHello).Methods("GET")

	// Static "/" must be last in code
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(port, router))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	var in []inMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&in)
	if err != nil {
		return
	}
	var out []outMsg
	for k, i := range in {
		fmt.Println(i, k)
		out = append(out, outMsg{Phone: i.Phone, Status: "ok"})
	}
	// fmt.Println(in.Phone)
	var buf []byte
	buf, err = json.Marshal(out)
	_, err = w.Write(buf)
}
