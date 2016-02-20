package views

// Define views here

import (
	. "github.com/r0fls/reinhardt/src/views"
)

// Default home
func Home(r reinhardt.Request) Response {
	return Render(r.Url, "home.html")
}
