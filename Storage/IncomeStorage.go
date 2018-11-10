package storage

import (
	"strconv"

	"github.com/geniuscasio/rest-gotrello/entities"
)

var incomes []entities.Income

//Save save entity to persistent
func Save(i entities.Income) {
	incomes = append(incomes, i)
}

//GetAll returns all entities in DB
func GetAll() []entities.Income {
	return incomes
}

//GetByID returns
func GetByID(id string) entities.Income {
	var _id int64
	if id == "" {
		_id = 1
	} else {
		_id, _ = strconv.ParseInt(id, 0, 64)
	}
	income := entities.Income{ID: _id, Amount: 50.0, Hint: "GetById result", Tags: nil}
	incomes = append(incomes, income)
	return income
}
