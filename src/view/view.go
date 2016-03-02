package view

import (
	//"fmt"
	"encoding/json"
	"github.com/r0fls/reinhardt/src/config"
	"github.com/r0fls/reinhardt/src/template"
	//"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func Render(temp string, r Request) {
	dir, _ := os.Getwd()
	f := strings.Join([]string{dir, "settings.json"}, "/")
	config := config.Load_config(f)
	f = strings.Join([]string{config.Home, config.Apps[0], config.Templates[0], temp}, "/")
	s := template.Sub{config.Static, r.Request.URL.Path}
	text, err := ioutil.ReadFile(f)
	template.Load(string(text), s, r.Response)
	check(err)
	//fmt.Fprintf(r.Response, string(text), html.EscapeString(r.Request.URL.Path))
}

func RenderJSON(m interface{}, r Request) {
	enc := json.NewEncoder(r.Response)
	err := enc.Encode(m)
	check(err)
}

type View func(Request)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
