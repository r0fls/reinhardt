package main

import (
	"encoding/json"
	"fmt"
	"github.com/r0fls/reinhardt/src/view"
	"github.com/r0fls/reinhardt/test/app"
	//"html"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func load_views(appname string) {
}

func load_models(appname string) {
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
	Urls := app.Urls()
	for _, url := range Urls {
		http.HandleFunc(url.Slug, func(w http.ResponseWriter, r *http.Request) {
			//s := []string{config.Home, config.Apps[0], config.Templates[0], "home.html"}
			url.View(view.Response{w}, view.Request{r})
		})
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
