package url

import (
	//"github.com/r0fls/reinhardt/src/view"
	"net/http"
)

type Url struct {
	Slug string
	View func(http.ResponseWriter, *http.Request)
	//View view.View
}
