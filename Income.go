package main

import (  
	"fmt"
	"IncomeTag"
)

var Incomes []Income

type Income struct {
	Id     string      `json:"id,omitempty"`
	Amount float32       `json:"amount,omitempty"`
	Hint   string      `json:"hint,omitempty"`
	Tags   []IncomeTag `json:"tags,omitempty"`
}

func NewIncome(amount float32, hint string, tags []IncomeTag) Income {
	i := Income{Id: "1", Amount: amount, Hint: hint, Tags: tags}
	Incomes = append(Incomes, i)
	return i
}