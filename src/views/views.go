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
