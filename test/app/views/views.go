package views

// Define views here

import (
	. "github.com/r0fls/reinhardt/src/views"
)

// Default home
func Home(r Request) Response {
	return Render(r, "home.html")
}
