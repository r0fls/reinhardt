package view

import (
	"fmt"
	//	"github.com/julienschmidt/httprouter"
	"github.com/r0fls/reinhardt/src/config"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func Render(template string, r Request) {
	dir, _ := os.Getwd()
	f := strings.Join([]string{dir, "settings.json"}, "/")
	config := config.Load_config(f)
	//text := respool.Read(strings.Join(append(base, template), "/"))
	f = strings.Join([]string{config.Home, config.Apps[0], config.Templates[0], template}, "/")
	text, err := ioutil.ReadFile(f)
	check(err)
	fmt.Fprintf(r.Response, string(text), html.EscapeString(r.Request.URL.Path))
	/*
		var text []byte
		var err error

		for i, _ := range config.Templates {
			print(strings.Join([]string{config.Home, config.Templates[i], template}, "/"), "\n")
			text, err = ioutil.ReadFile(strings.Join([]string{config.Home, config.Templates[i], template}, "/"))
			if err == nil {
				fmt.Fprintf(r.Response, string(text), html.EscapeString(r.Request.URL.Path))
				break
			}
			print(string(text))
		}
	*/
}

type View func(template string, base []string) func(w http.ResponseWriter, r *http.Request)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
