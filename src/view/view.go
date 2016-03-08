package view

import (
	"encoding/json"
	"github.com/r0fls/reinhardt/src/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type Request struct {
	Request  *http.Request
	Response http.ResponseWriter
}

type TemplateConfig struct {
	Config  config.Config
	Request Request
}

func (t TemplateConfig) Load(text string, name ...string) {
	if len(name) == 0 {
		name = append(name, "template")
	}
	tmpl, err := template.New(name[0]).Parse(text)
	check(err)
	err = tmpl.Execute(t.Request.Response, t)
	check(err)
}

func Render(temp string, r Request) {
	dir, _ := os.Getwd()
	f := strings.Join([]string{dir, "settings.json"}, "/")
	config := config.Load_config(f)
	f = strings.Join([]string{config.Home, config.Apps[0], config.Templates[0], temp}, "/")
	t := TemplateConfig{config, r}
	text, err := ioutil.ReadFile(f)
	t.Load(string(text))
	check(err)
}

func RenderJSON(m interface{}, r Request) {
	enc := json.NewEncoder(r.Response)
	err := enc.Encode(m)
	r.Response.Header().Set("Content-Type", "application/json")
	check(err)
}

type View func(Request)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
