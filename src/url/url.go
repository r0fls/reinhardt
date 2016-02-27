package url

import (
	//"github.com/r0fls/reinhardt/src/view"
	"net/http"
	"regexp"
)

type Url struct {
	Slug string
	View func(http.ResponseWriter, *http.Request)
	//View view.View
}

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler struct {
	routes []*route
}

func New(urls []Url) RegexpHandler {
	r := make([]*route, len(urls))
	return RegexpHandler{r}
}

func (h *RegexpHandler) handler(pattern *regexp.Regexp, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

func (h *RegexpHandler) Handler(pattern string, handler http.Handler) {
	re, _ := regexp.Compile(pattern)
	h.routes = append(h.routes, &route{re, handler})
}

func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.NotFound(w, r)
}
