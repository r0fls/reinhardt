package view

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

type Request struct {
	R *http.Request
}

type Response struct {
	W http.ResponseWriter
}

func Render(res Response, req Request, template string) {
	text, err := ioutil.ReadFile(template)
	if err != nil {
		fmt.Fprintf(res.W, string(text), html.EscapeString(req.R.URL.Path))
	} else {
		fmt.Fprintf(res.W, "error retrieving template %s 404", template)
	}
}

type View func(context Request, template string) Response

func response(e error) string {
	if e != nil {
		return "404"
	} else {
		return "200"
	}
}
