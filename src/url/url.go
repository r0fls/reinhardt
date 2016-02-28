package url

import (
	//"github.com/r0fls/reinhardt/src/view"
	"fmt"
	"net/http"
	"regexp"
)

type Url struct {
	Slug string
	View func(http.ResponseWriter, *http.Request)
	Re   *regexp.Regexp
	//View view.View
}

func URL(pattern string, v func(http.ResponseWriter, *http.Request)) Url {
	re := regexp.MustCompile(pattern)
	return Url{pattern, v, re}
}

func (u Url) Route(w http.ResponseWriter, r *http.Request) {
	re, err := regexp.Compile(u.Slug)
	print(err, "\n")
	print(r.URL.Path, "\n")
	if re.MatchString(r.URL.Path) {
		u.View(w, r)
	} else {
		http.NotFound(w, r)
	}
}

type Urls []Url

func (urls Urls) Routes(w http.ResponseWriter, r *http.Request) {
	for _, u := range urls {
		if u.Re.MatchString(r.URL.Path) {
			u.View(w, r)
			return
		}
	}

	w.Write([]byte(fmt.Sprintf("Did not match:%s", r.URL.Path)))
}
