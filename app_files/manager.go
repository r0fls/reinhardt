package manager

import (
	"html"
	"log"
	"net/http"
)

func load_views(appname string) {
}

func load_models(appname string) {
}

func load_app(appname string) {
}

func run_server() {
	//http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	// for app in conf
	// load_app
}

if os.Args[1] == "runserver" {
	run_server()
}
