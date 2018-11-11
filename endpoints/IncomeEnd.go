package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	storage "github.com/geniuscasio/rest-gotrello/Storage"
	"github.com/geniuscasio/rest-gotrello/entities"
	"github.com/gorilla/mux"
)

//Create entity
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create invoked!")
	var income entities.Income
	_ = json.NewDecoder(r.Body).Decode(&income)
	storage.Save(income)
	json.NewEncoder(w).Encode(income)
}

//Get entity
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Get")
	params := mux.Vars(r)
	_id := params["id"]
	if _id == "" {
		fmt.Println("id empty")
		json.NewEncoder(w).Encode(storage.GetAll())
	} else {
		json.NewEncoder(w).Encode(storage.GetByID(params["id"]))
	}
}

//Update entity
func Update(w http.ResponseWriter, r *http.Request) {

}

//Delete entity
func Delete(w http.ResponseWriter, r *http.Request) {

}

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
