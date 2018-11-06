package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"fmt"
)

type Person struct {
	ID        string   `json:"id,omitempty`
	Firstname string   `json:"firstname,omitempty`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
	json.NewEncoder(w).Encode(people)
}
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func db_prepare() {
	host := os.Getenv("db_host") 
	port := os.Getenv("db_port") 
	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	name := os.Getenv("db_name")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s"+
		"password=%s dbname=%s sslmode=%s sslmode=disable",
		host, port, user, password, name)

	fmt.Println(psqlInfo);
}

func main() {
	router := mux.NewRouter()
	db_prepare()
	port := os.Getenv("PORT")

	fmt.Println("Message")
	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		port = ":"+port
	}

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")	
	router.HandleFunc("/people/{id}", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeopleEndpoint).Methods("DELETE")
	// // Income sources tags
	// router.HandleFunc("/incomeTag", GetIncomeTagEndpoint).Methods("GET")	
	// router.HandleFunc("/incomeTag/{id}", GetIncomeTagEndpoint).Methods("GET")
	// router.HandleFunc("/incomeTag/{id}", CreateIncomeTagEndpoint).Methods("POST")
	// router.HandleFunc("/incomeTag/{id}", DeleteIncomeTagEndpoint).Methods("DELETE")
	// // Outcome sources tags
	// router.HandleFunc("/outcomeTag", GetOutcomeTagEndpoint).Methods("GET")
	// router.HandleFunc("/outcomeTag/{id}", GetOutcomeTagEndpoint).Methods("GET")
	// router.HandleFunc("/outcomeTag/{id}", CreateOutcomeTagEndpoint).Methods("POST")
	// router.HandleFunc("/outcomeTag/{id}", DeleteOutcomeTagEndpoint).Methods("DELETE")
	// // Route Income
	// router.HandleFunc("/income", GetIncomeEndpoint).Methods("GET")	
	// router.HandleFunc("/income/{id}", GetIncomeEndpoint).Methods("GET")
	// router.HandleFunc("/income/{id}", CreateIncomeEndpoint).Methods("POST")
	// router.HandleFunc("/income/{id}", DeleteIncomeEndpoint).Methods("DELETE")
	// // Route Outcome
	// router.HandleFunc("/outcome", GetOutcomeEndpoint).Methods("GET")
	// router.HandleFunc("/outcome/{id}", GetOutcomeEndpoint).Methods("GET")
	// router.HandleFunc("/outcome/{id}", CreateOutcomeEndpoint).Methods("POST")
	// router.HandleFunc("/outcome/{id}", DeleteOutcomeEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(port, router))
}