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

// type Person struct {
// 	ID        string   `json:"id,omitempty"`
// 	Firstname string   `json:"firstname,omitempty"`
// 	Lastname  string   `json:"lastname,omitempty"`
// 	Address   *Address `json:"address,omitempty"`
// }
// type Address struct {
// 	City  string `json:"city,omitempty"`
// 	State string `json:"state,omitempty"`
// }

// var people []Person

//GetPersonEndpoint comment
// func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for _, item := range people {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Person{})
// }
// func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	for key, value := range r.Form {
// 		fmt.Printf("%s = %s\n", key, value)
// 	}
// 	json.NewEncoder(w).Encode(people)
// }
// func CreatePeopleEndpoint(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var person Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	person.ID = params["id"]
// 	people = append(people, person)
// 	json.NewEncoder(w).Encode(people)
// }
// func DeletePeopleEndpoint(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for index, item := range people {
// 		if item.ID == params["id"] {
// 			people = append(people[:index], people[index+1:]...)
// 			break
// 		}
// 		json.NewEncoder(w).Encode(people)
// 	}
// }

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
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	router.HandleFunc("/", endpoints.Get).Methods("GET")
	router.HandleFunc("/income/{id}", endpoints.Get).Methods("GET")
	router.HandleFunc("income/", endpoints.Create).Methods("POST")
	router.HandleFunc("/income/", endpoints.Get).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}
