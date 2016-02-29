package views

// Define views here

import (
	"github.com/r0fls/reinhardt/src/view"
	"net/http"
)

// Default home
func Home(r view.Request) {
	view.Render("home.html", r)
}
