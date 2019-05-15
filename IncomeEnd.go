package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	storage "github.com/geniuscasio/rest-gotrello/Storage"
	"github.com/geniuscasio/rest-gotrello/entities"
)

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
		fmt.Println("error decoding income", err.Error())
	}
	storage.Save(income)
	created, _ := json.Marshal(income)
	w.WriteHeader(201)
	w.Write(created)
}

//Get entity
func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Get")
	userID := ""
	sessionID, err := r.Cookie("sessionId")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		userID = GetUserBySession(sessionID.Value).Login
	}
	data, _ := json.Marshal(getIncome("", userID))
	w.Write(data)
}
