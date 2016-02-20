package main

import (
	"encoding/json"
	"fmt"
	"html"
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
	Address string
	Port    string
	Users   []string
	Groups  []string
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
	//http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	ap := []string{config.Address, config.Port}
	go log.Fatal(http.ListenAndServe(strings.Join(ap, ":"), nil))
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
