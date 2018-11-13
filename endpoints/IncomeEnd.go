package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	storage "github.com/geniuscasio/rest-gotrello/storage"
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
