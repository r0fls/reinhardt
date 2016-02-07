package main

import (
	//"fmt"
	//"html"
	//"log"
	//"net/http"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func new_project(name string) {
	m := []string{name, "/app", "/models"}
	os.MkdirAll(strings.Join(m, ""), 0700)
	mf := []string{name, "/app", "/models", "/models.go"}
	os.Create(strings.Join(mf, ""))
	v := []string{name, "/app", "/views"}
	os.Mkdir(strings.Join(v, ""), 0700)
	vf := []string{name, "/app", "/views", "/views.go"}
	os.Create(strings.Join(vf, ""))
	t := []string{name, "/app", "/temps"}
	os.Mkdir(strings.Join(t, ""), 0700)
	s := []string{name, "/settings.toml"}
	settings, err := ioutil.ReadFile("settings.toml")
	check(err)
	err = ioutil.WriteFile(strings.Join(s, ""), settings, 0644)
	check(err)
	u := []string{name, "/urls.go"}
	os.Create(strings.Join(u, ""))
}

// move to a mangager in project?
func load_views(appname string) {
}

func load_models(appname string) {
}

func load_app(appname string) {
}

func run_server() {
	// for app in conf
	// load_app
}
func main() {
	if os.Args[1] == "newproject" {
		new_project(os.Args[2])
	}
	if os.Args[1] == "runserver" {
		run_server()
	}
	//example http server
	/*
		//http.Handle("/foo", fooHandler)
		http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}
