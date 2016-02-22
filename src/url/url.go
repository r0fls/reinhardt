package url

import (
	"github.com/r0fls/reinhardt/src/view"
)

type Url struct {
	Slug string
	View func(view.Response, view.Request, []string)
}
