package views

// Define views here

import (
	"github.com/r0fls/reinhardt/src/view"
	"net/http"
)

// Default home
func Home(w http.ResponseWriter, r *http.Request) {
	view.Render("home.html", view.Request{r, w})
}
