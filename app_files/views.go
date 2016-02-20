package views

// Define views here

import (
	. "github.com/r0fls/reinhardt/src/view"
)

// Default home
View Home(r Request) Response {
	return Render(r, "home.html")
}
