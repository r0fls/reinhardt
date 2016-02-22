package views

// Define views here

import (
	"github.com/r0fls/reinhardt/src/view"
	"net/http"
)

// Default home
func Home(base []string) func(w http.ResponseWriter, r *http.Request) {
	return view.Render("home.html", base)
}
