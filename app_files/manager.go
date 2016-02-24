package main

import (
	"encoding/json"
	"fmt"
	"github.com/r0fls/reinhardt/test/app"
	"log"
	"net/http"
	"os"
	"strings"
)

func load_views(appname string) {
}

func load_models(appname string) {
	//model.Connect("postgres", "postgres", "localhost", "235711", "test")
}

func load_app(appname string) {
	//http.Handle("/foo", fooHandler)
}

type Config struct {
	Address   string
	Port      string
	Home      string
	Templates []string
	Apps      []string
	DB        DBConfig
}

type DBConfig struct {
	Type string
	User string
	Name string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func load_config(location string) Config {
	file, _ := os.Open(location)
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	return config
}

func run_server(location string) {
	config := load_config(location)
	// should loop through all apps
	base := []string{config.Home, config.Apps[0], config.Templates[0]}
	Urls := app.Urls()
	for i, _ := range Urls {
		http.Handle(Urls[i].Slug, http.HandlerFunc(Urls[i].View(base)))
	}
	ap := []string{config.Address, config.Port}
	log.Fatal(http.ListenAndServe(strings.Join(ap, ":"), nil))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires more args")
	} else if os.Args[1] == "runserver" {
		run_server("settings.json")
	}
}
