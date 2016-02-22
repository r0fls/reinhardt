package view

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	R *http.Request
}

type Response struct {
	W http.ResponseWriter
}

func Render(template string, base []string) func(w http.ResponseWriter, r *http.Request) {
	text, err := ioutil.ReadFile(strings.Join(append(base, template), "/"))
	check(err)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(text), html.EscapeString(r.URL.Path))
	}
}

type View func(context Request, template string) Response

func check(e error) {
	if e != nil {
		panic(e)
	}
}
