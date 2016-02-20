package views

type Request struct {
	Headers    string
	Parameters string
	Url        string
}

type Response struct {
	Status string
	Body   string
}

func Render(context Request, template string) r Response {

	s := []string{name, "app", "views", "views.go"}
	text, err := ioutil.ReadFile("app_files/views.go")
	res := response(err)
	return Response{res, text)
}

func response(e error) string {
	if e != nil {
		return "404"
	} else {
		return "200"
	}
}
