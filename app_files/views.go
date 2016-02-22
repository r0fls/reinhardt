package views

// Define views here

import (
	"github.com/r0fls/reinhardt/src/view"
)

// Default home
func Home(res view.Response, r view.Request, base []string) {
	view.Render(res, r, "home.html", base)
}
