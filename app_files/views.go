package views

// Define views here

import (
	"github.com/r0fls/reinhardt/src/view"
)

// Default home
func Home(r view.Request) {
	view.Render("home.html", r)
}
