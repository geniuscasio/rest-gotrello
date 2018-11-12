package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"models"
	"fmt"
)
var IncomeStorage []Income
func GetAllIncomeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range IncomeStorage {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Income{})
}
func GetIncomeEndpoint(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
	json.NewEncoder(w).Encode(Income)
}
func CreateIncomeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	income := Income{}
	_ = json.NewDecoder(r.Body).Decode(&income)
	income.ID = params["id"]
	IncomeStorage = append(IncomeStorage, income)
	json.NewEncoder(w).Encode(IncomeStorage)
}
func DeleteIncomeEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range IncomeStorage {
		if item.ID == params["id"] {
			Income = append(IncomeStorage[:index], IncomeStorage[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(IncomeStorage)
	}
}