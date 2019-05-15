package storage

import (
	"strconv"
	"time"

	"github.com/geniuscasio/rest-gotrello/entities"
	_ "github.com/lib/pq"
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

//Get returns Income with given Id (always return something for now)
func Get(id string) []entities.Income {
	var _id int64
	if id == "" {
		_id = 1
	} else {
		_id, _ = strconv.ParseInt(id, 0, 64)
	}
	tag1 := entities.IncomeTag{ID: 1, Name: "Tag1", Description: "abc", Aliases: []string{"a"}}
	tag2 := entities.IncomeTag{ID: 2, Name: "Tag2", Description: "abc", Aliases: []string{"a"}}
	tag3 := entities.IncomeTag{ID: 3, Name: "Tag3", Description: "abc", Aliases: []string{"a"}}
	tt := []entities.IncomeTag{tag1, tag2, tag3}
	income := entities.Income{ID: _id, Amount: 50.0, Date: time.Now(), Hint: "GetById result", Tags: tt}
	incomes = append(incomes, income)
	return incomes
}
