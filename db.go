package main

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

const _SQLInsertUser = "INSERT INTO my_users(username, password) VALUES (?, ?)"

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
	_, r = c.Exec(`
		INSERT INTO MY_USERS(username, password) select "admin" as "username", "-995833633" as "password"
	`)
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
	stmt, err := getDB().Prepare("SELECT password FROM MY_USERS WHERE username=$1")
	if err != nil {
		fmt.Errorf(err.Error())
	}
	res, _ := stmt.Query(name)
	for res.Next() {
		err = res.Scan(&pass)
	}
	return ""
}
