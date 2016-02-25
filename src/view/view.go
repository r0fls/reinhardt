package view

import (
	"fmt"
	//	"github.com/julienschmidt/httprouter"
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
	//text := respool.Read(strings.Join(append(base, template), "/"))
	text, err := ioutil.ReadFile(strings.Join(append(base, template), "/"))
	check(err)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(text), html.EscapeString(r.URL.Path))
	}
}

type View func(template string, base []string) func(w http.ResponseWriter, r *http.Request)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
