package view

import (
	"io/ioutil"
)

type Request struct {
	Headers    string
	Parameters string
	Url        string
}

type Response struct {
	Status string
	Body   []byte
}

func Render(context Request, template string) Response {
	// this should be the template, inside the templates dir from config
	//s := []string{template, "app", "views", "views.go"}
	text, err := ioutil.ReadFile("app_files/views.go")
	res := response(err)
	return Response{res, text}
}

type View func(context Request, template string) Response

func response(e error) string {
	if e != nil {
		return "404"
	} else {
		return "200"
	}
}
