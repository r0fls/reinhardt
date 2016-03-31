package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/r0fls/reinhardt/src/config"
	"log"
	"os"
	"strings"
)

// TODO
// 0. API should have commands that are independant of the database type
// 1. create database from models
//    a. Handle optional args (unique, etc...)
//    b. constraints (using func or args in initializer funcs)
// 2. create functions that allow access from views
//    a. insert (needs to be sanitized)
//    b. update (needs to be sanitized)
//    c. delete (needs to be sanitized)
//    d. get (needs to be sanitized)
//    e. filter should be able to be used with some of the above, e.g. delete

// Takes a db type (eg. postgres) username and dbname (should take pswd too)
// This should only be called once, and reused for all connections after (leave
// connections open)

type connection struct {
	DB     *sql.DB
	dbtype string
}

type ModelType struct {
	connection
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

func configConnect(config config.Config) *sql.DB {
	return Connect(config.DB.Type, config.DB.User, config.DB.Name, config.DB.IP, config.DB.Pass)
}

func (m Model) AddModel(name string) {
	dir, _ := os.Getwd()
	f := strings.Join([]string{dir, "settings.json"}, "/")
	config := config.Load_config(f)
	db := configConnect(config)
	mt := ModelType{connection{db, config.DB.Type}, name, []Field{}}
	m[name] = &mt
}

func NewModel(name string) Model {
	m := make(map[string]*ModelType)
	dir, _ := os.Getwd()
	f := strings.Join([]string{dir, "settings.json"}, "/")
	config := config.Load_config(f)
	db := configConnect(config)
	mt := ModelType{connection{db, config.DB.Type}, name, []Field{}}
	m[name] = &mt
	return m
}

// These are all for postgres right now -- should use
// DB type from

// Should implement the following types with their spelling, as
// they are specified by SQL (http://www.postgresql.org/docs/9.4/static/datatype.html):
// bigint, bit, bit varying, boolean, char, character varying, character, varchar,
// date, double precision, integer, interval, numeric, decimal, real, smallint, time
// (with or without time zone), timestamp (with or without time zone), xml.

// this should accept additional arguments (default, etc..)
func (m *ModelType) IntegerField(args ...string) {
	m.F = append(m.F, Field{args[0], "integer"})
}

func (m *ModelType) CharacterField(args ...string) {
	m.F = append(m.F, Field{args[0], "varchar(80)"})
}

// should implement float8 as an option
func (m *ModelType) RealField(args ...string) {
	m.F = append(m.F, Field{args[0], "float4"})
}

func (m *ModelType) TextField(args ...string) {
	m.F = append(m.F, Field{args[0], "text"})
}

func (m *ModelType) DateField(args ...string) {
	m.F = append(m.F, Field{args[0], "date"})
}

func (m *ModelType) TimeField(args ...string) {
	m.F = append(m.F, Field{args[0], "time"})
}

func (m *ModelType) TimestampField(args ...string) {
	m.F = append(m.F, Field{args[0], "timestamp"})
}

func (m *ModelType) JsonField(args ...string) {
	m.F = append(m.F, Field{args[0], "json"})
}

func (m *ModelType) JsonBField(args ...string) {
	m.F = append(m.F, Field{args[0], "bjson"})
}

func (m *ModelType) ByteField(args ...string) {
	m.F = append(m.F, Field{args[0], "bytea"})
}

func (m *ModelType) BooleanField(args ...string) {
	m.F = append(m.F, Field{args[0], "bool"})
}

func (m *ModelType) XmlField(args ...string) {
	m.F = append(m.F, Field{args[0], "xml"})
}

func (m *ModelType) UuidField(args ...string) {
	m.F = append(m.F, Field{args[0], "uuid"})
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

func CreateTable(m ModelType) {
	print("Creating Models", "\n")
	for i, _ := range m.F {
		if i == 0 {
			s := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s %s);", m.Name, m.F[i].Name, m.F[i].Type)
			_, err := m.DB.Query(s)
			fmt.Println(s)
			check(err)
		} else {
			s := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;",
				m.Name, m.F[i].Name, m.F[i].Type)
			_, err := m.DB.Query(s)
			fmt.Println(s)
			check(err)
		}
	}
}

func (m Model) CreateTables() {
	for _, model := range m {
		CreateTable(*model)
	}
}

func (m *ModelType) Insert(v []string) {
	s := fmt.Sprintf("INSERT INTO %s(%s) VALUES ('%s');",
		m.Name, strings.Join(m.Cols(), ", "),
		strings.Join(v, "', '"))
	fmt.Println(s)
	_, err := m.DB.Query(s)
	check(err)
}

func (m Model) Insert(model string, v []string) {
	s := fmt.Sprintf("INSERT INTO %s(%s) VALUES ('%s');",
		m[model].Name, strings.Join(m[model].Cols(), ", "),
		strings.Join(v, "', '"))
	fmt.Println(s)
	_, err := m[model].DB.Query(s)
	check(err)
}

func get(db *sql.DB, column string, table string, filter string) *sql.Rows {
	s := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", column, table, column)
	rows, err := db.Query(s, filter)
	check(err)
	return rows
}
