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

// Takes a db type (eg. postgres) username and dbname (should take pswd too)
// This should only be called once, and reused for all connections after (leave
// connections open)
func Connect(dbtype string, username string, dbname string,
	ip string, password string) *sql.DB {

	s := fmt.Sprintf("%s://%s:%s@%s/%s",
		dbtype, username, password, ip, dbname)
	db, err := sql.Open(dbtype, fmt.Sprintf(s))
	check(err)
	return db
}

type Model struct {
	Name string
	F    []Field
}

type Field struct {
	Name string
	T    string
}

type Type struct {
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func CreateTable(db *sql.DB, m Model) {
	// needs to loop through all fields
	s := fmt.Sprintf(`CREATE TABLE %s (%s %s);`, m.Name, m.F[0].Name, m.F[0].T)
	_, err := db.Query(s)
	check(err)
}

// this needs to be weary of SQL injection
func Insert(db *sql.DB, m Model) {
	// needs to loop through all fields
	s := fmt.Sprintf(`INSERT INTO %s VALUES (%s);`, m.Name, m.F[0].Name)
	_, err := db.Query(s)
	check(err)
}

func get(db *sql.DB, column string, table string, filter string) *sql.Rows {
	s := fmt.Sprintf("SELECT %s FROM %s WHERE age=?", table, column)
	rows, err := db.Query(s, filter)
	check(err)
	return rows
}
