package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
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
	Templates []string
	Users     []string
	Groups    []string
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
	s := []string{config.Templates[0], "home.html"}
	text, err := ioutil.ReadFile(strings.Join(s, "/"))
	check(err)
	//http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, string(text), html.EscapeString(r.URL.Path))
	})
	ap := []string{config.Address, config.Port}
	log.Fatal(http.ListenAndServe(strings.Join(ap, ":"), nil))
	// for app in conf
	// load_app
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires more args")
	} else if os.Args[1] == "runserver" {
		run_server("settings.json")
	}
}
