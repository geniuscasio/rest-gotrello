package storage

import (
	"strconv"
	"time"

	"github.com/geniuscasio/rest-gotrello/entities"
)

var incomes []entities.Income

//TODO: Storage will be working with PostgreSQL

//Save save entity to storage
func Save(i entities.Income) {
	incomes = append(incomes, i)
}

//GetAll returns all entities in storage
func GetAll() []entities.Income {
	return incomes
}

//GetByID returns Income with given Id (always return something for now)
func GetByID(id string) entities.Income {
	var _id int64
	if id == "" {
		_id = 1
	} else {
		_id, _ = strconv.ParseInt(id, 0, 64)
	}
	income := entities.Income{ID: _id, Amount: 50.0, Date: time.Now(), Hint: "GetById result", Tags: nil}
	incomes = append(incomes, income)
	return income
}
