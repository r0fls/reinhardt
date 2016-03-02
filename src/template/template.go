package template

import (
	"log"
	"net/http"
	"text/template"
)

type Sub struct {
	Static string
	Path   string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Load(text string, s Sub, w http.ResponseWriter, name ...string) {
	if len(name) == 0 {
		name = append(name, "template")
	}
	tmpl, err := template.New(name[0]).Parse(text)
	check(err)
	err = tmpl.Execute(w, s)
	check(err)
}
