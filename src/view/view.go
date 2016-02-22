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

func Render(res Response, req Request, template string, base []string) {
	text, err := ioutil.ReadFile(strings.Join(append(base, template), "/"))
	check(err)
	fmt.Fprintf(res.W, string(text), html.EscapeString(req.R.URL.Path))
}

type View func(context Request, template string) Response

func check(e error) {
	if e != nil {
		panic(e)
	}
}
