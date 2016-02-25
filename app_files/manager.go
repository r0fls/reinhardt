package main

import (
	"encoding/json"
	"fmt"
	"github.com/r0fls/reinhardt/src/config"
	"github.com/r0fls/reinhardt/test/app"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func load_views(appname string) {
}

func load_models(appname string) {
	//model.Connect("postgres", "postgres", "test", "localhost", "235711")
	c := config.Load_config("settings.json")
	print(c.DB.Name, "\n")
	print(c.DB.Type, "\n")
	print(c.DB.IP, "\n")
	print(c.DB.Pass, "\n")
	print(c.DB.User, "\n")

}

func load_app(appname string) {
	//http.Handle("/foo", fooHandler)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func run_server(location string) {
	config := config.Load_config(location)
	// should loop through all apps
	base := []string{config.Home, config.Apps[0], config.Templates[0]}
	Urls := app.Urls()
	for i, _ := range Urls {
		http.Handle(Urls[i].Slug, http.HandlerFunc(Urls[i].View(base)))
	}
	ap := []string{config.Address, config.Port}
	log.Fatal(http.ListenAndServe(strings.Join(ap, ":"), nil))
}

func init() {
	config := config.Load_config("settings.json")
	s := []string{"app", "models", ".settings.json"}
	fileJson, _ := json.Marshal(config.DB)
	err := ioutil.WriteFile(strings.Join(s, "/"), fileJson, 0644)
	check(err)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires more args")
	} else if os.Args[1] == "runserver" {
		run_server("settings.json")
	} else if os.Args[1] == "syncdb" {
		load_models("settings.json")
	}
}
