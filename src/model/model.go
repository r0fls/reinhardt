package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
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

type ModelType struct {
	Name string
	F    []Field
}

type Model map[string]*ModelType

type Field struct {
	Name string
	Type string
}

type Fields []Field

type Value struct {
	Cols []string
	Row  []string
}

func Connect(dbtype string, username string, dbname string,
	ip string, password string) *sql.DB {

	s := fmt.Sprintf("%s://%s:%s@%s/%s",
		dbtype, username, password, ip, dbname)
	db, err := sql.Open(dbtype, fmt.Sprintf(s))
	check(err)
	return db
}

func (m Model) AddModel(name string) {
	mt := ModelType{name, []Field{}}
	m[name] = &mt
}

func NewModel(name string) Model {
	m := make(map[string]*ModelType)
	mt := ModelType{name, []Field{}}
	m[name] = &mt
	return m
}

func (m *ModelType) IntegerField(name string) {
	m.F = append(m.F, Field{name, "integer"})
}

func (m *ModelType) CharacterField(name string) {
	m.F = append(m.F, Field{name, "varchar(80)"})
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (m ModelType) Cols() []string {
	var cols []string
	for _, field := range m.F {
		cols = append(cols, field.Name)
	}
	return cols
}

func CreateTable(db *sql.DB, m ModelType) {
	// needs to loop through all fields
	print("Creating Models", "\n")
	for i, _ := range m.F {
		if i == 0 {
			s := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s %s);", m.Name, m.F[i].Name, m.F[i].Type)
			_, err := db.Query(s)
			fmt.Println(s)
			check(err)
		} else {
			s := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;",
				m.Name, m.F[i].Name, m.F[i].Type)
			_, err := db.Query(s)
			fmt.Println(s)
			check(err)
		}
	}
}

func CreateTables(db *sql.DB, m Model) {
	for _, model := range m {
		CreateTable(db, *model)
	}
}

// this needs a custom formatting function:
// if string - wrap in ''; int - nowrap, etc..
func (m *ModelType) Insert(db *sql.DB, v Value) {
	s := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s);",
		m.Name, strings.Join(v.Cols, ", "),
		strings.Join(v.Row, " "))
	_, err := db.Query(s)
	check(err)
}

func get(db *sql.DB, column string, table string, filter string) *sql.Rows {
	s := fmt.Sprintf("SELECT %s FROM %s WHERE age=?", table, column)
	rows, err := db.Query(s, filter)
	check(err)
	return rows
}
