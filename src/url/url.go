package url

import (
	"net/http"
)

type Url struct {
	Slug string
	View func([]string) func(w http.ResponseWriter, r *http.Request)
}
