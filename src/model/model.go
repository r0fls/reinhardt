package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// TODO
// 1. connect function
//    -- uses settings
// 2. create database from models
//    FIELDS
//    a. Int
//    b. Float
//    c. String
//    d. Bool
//    e. Datetime
// 3. create functions that allow access from views
//    a. insert
//    b. update
//    c. delete
//    d. get
//    e. filter (should be able to be used with some of the above, e.g. delete)

// takes a db type (eg. postgres) username and dbname (should take pswd too)
func Connect(dbtype string, username string, dbname string) *sql.DB {
	db, err := sql.Open(dbtype, fmt.Sprintf("user=%s dbname=%s sslmode=verify-full", username, dbname))
	check(err)
	return db
}

type Model struct {
	Fields []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func CreateTables(db *sql.DB, m []Model) {
}

func Select(db *sql.DB, table_name string, filter string) *sql.Rows {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table_name))
	check(err)
	return rows
}
