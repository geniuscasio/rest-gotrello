package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	storage "github.com/geniuscasio/rest-gotrello/Storage"
	"github.com/geniuscasio/rest-gotrello/entities"
	"github.com/gorilla/mux"
)

type SecureHandler func(r *http.Request, rb *ResponseBox)

//Create entity
func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create invoked!")
	r.ParseForm()
	fmt.Println(r.Form)
	defer r.Body.Close()
	fmt.Println(r.Body)
	var income entities.Income
	err := json.NewDecoder(r.Body).Decode(&income)
	if err != nil {
		fmt.Println(err.Error())
	}
	storage.Save(income)
	created, _ := json.Marshal(income)
	w.WriteHeader(201)
	w.Write(created)
}

//Get entity
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Get")
	params := mux.Vars(r)
	_id := params["id"]
	fmt.Print(_id)
	if _id == "" {
		data, _ := json.Marshal(storage.GetAll())
		w.Write(data)
	} else {
		data, _ := json.Marshal(storage.GetByID(params["id"]))
		w.Write(data)
	}
}

//Update entity
func Update(r *http.Request, rb *ResponseBox) {

}

//Delete entity
func Delete(r *http.Request, rb *ResponseBox) {

}
