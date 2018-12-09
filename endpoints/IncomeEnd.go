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
func Create(r *http.Request, rb *ResponseBox) {
	fmt.Println("Create invoked!")
	r.ParseForm()
	fmt.Println(r.Form)
	defer r.Body.Close()
	/*
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//handle read response error
			fmt.Println("Error")
		}
		fmt.Printf("%s\n", string(body))*/

	fmt.Println(r.Body)
	var income entities.Income
	err := json.NewDecoder(r.Body).Decode(&income)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(income.Hint)
	storage.Save(income)
	rb.Content = income
}

//Get entity
func Get(r *http.Request, rb *ResponseBox) {
	fmt.Println("Inside Get")
	params := mux.Vars(r)
	_id := params["id"]
	fmt.Print(_id)
	if _id == "" {
		rb.Content = storage.GetAll()
	} else {
		rb.Content = storage.GetByID(params["id"])
	}
}

//Update entity
func Update(r *http.Request, rb *ResponseBox) {

}

//Delete entity
func Delete(r *http.Request, rb *ResponseBox) {

}
