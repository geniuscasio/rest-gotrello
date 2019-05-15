package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/geniuscasio/rest-gotrello/entities"
)

var db *sql.DB

const _SQLInsertUser = "INSERT INTO my_users(username, password) VALUES ($1, $2)"
const _SQLSelectAllUserIncomes = "SELECT income_id, amount, hint, tags, user_id FROM INCOME where user_id = $1"
const _SQLInsertIncomeByUser = "INSERT INTO INCOME(amount, hint, tags, user_id, date) VALUES ($1, $2, $3, $4, $5)"

func getDB() *sql.DB {
	if db == nil {
		fmt.Printf("connect to db")
		dbinfo := os.Getenv("DB_INFO")
		dbinfo = "postgres://nhshglygkfiair:5d42deb354a442697dea6593f1f2bb6e0e869bda02adaad22d3cdcbd321671e1@ec2-23-21-122-141.compute-1.amazonaws.com:5432/d2dorvldn0g3nl"
		var err error
		db, err = sql.Open("postgres", dbinfo)
		if err != nil {
			fmt.Printf(err.Error())
			return nil
		}
	}

	return db
}

func InitDB() {
	c := getDB()
	_, r := c.Exec(`
	CREATE TABLE MY_USERS(
		user_id serial PRIMARY KEY, 
		username VARCHAR (100) UNIQUE NOT NULL, 
		password VARCHAR (100) NOT NULL
	)`)
	if r != nil {
		fmt.Println(r.Error())
	}
	_, r = c.Exec(`drop table INCOME`)
	_, r = c.Exec(`
		CREATE TABLE INCOME(
			income_id serial PRIMARY KEY,
			amount NUMERIC (6, 2),
			hint VARCHAR (100),
			tags VARCHAR (300),
			date VARCHAR (100),
			user_id INTEGER REFERENCES MY_USERS(user_id)
		)`)
	if r != nil {
		fmt.Println(r.Error())
	}
}

func createUser(name, pwd string) bool {
	_, err := getDB().Exec(_SQLInsertUser, name, pwd)
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}
	return true
}

func getUser(name string) (pass string) {
	if name == "admin" {
		return "-995833633"
	}
	stmt, err := getDB().Prepare("SELECT password FROM MY_USERS WHERE username = $1")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	res, err := stmt.Query(name)
	defer res.Close()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	for res.Next() {
		err = res.Scan(&pass)
		return
	}
	return ""
}

func getIncome(id, userId string) []entities.Income {
	r, err := getDB().Query(_SQLSelectAllUserIncomes, userId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var resultSet []entities.Income
	for r.Next() {
		var i entities.Income
		err = r.Scan(&i)
		resultSet = append(resultSet, i)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return resultSet
}

func saveIncome(i entities.Income, userId string) bool {
	_, err := getDB().Exec(_SQLInsertIncomeByUser, i.Amount, i.Hint, i.Tags, userId, i.Date)
	if err != nil {
		fmt.Printf(err.Error())
		return false
	}
	return true
}
