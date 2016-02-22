package views

// Define views here

import (
	"github.com/r0fls/reinhardt/src/view"
)

// Default home
func Home(res view.Response, r view.Request) {
	view.Render(res, r, "home.html")
}
