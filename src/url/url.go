package url

import (
	"fmt"
	"github.com/r0fls/reinhardt/src/view"
	"net/http"
	"regexp"
)

type Url struct {
	Slug string
	View func(view.Request)
	Re   *regexp.Regexp
	//View view.View
}

func URL(pattern string, v func(view.Request)) Url {
	re := regexp.MustCompile(pattern)
	return Url{pattern, v, re}
}

type Urls []Url

func (urls Urls) Routes(w http.ResponseWriter, r *http.Request) {
	for _, u := range urls {
		if u.Re.MatchString(r.URL.Path) {
			u.View(view.Request{r, w})
			return
		}
	}

	w.Write([]byte(fmt.Sprintf("Did not match:%s", r.URL.Path)))
}
