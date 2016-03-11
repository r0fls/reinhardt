package app

import (
	"github.com/r0fls/reinhardt/src/url"
	"{{.Local}}/{{index .Apps 0}}/views"
)

func Urls() url.Urls {
	return []url.Url{
		url.URL("/bar", views.Home),
		url.URL("/$", views.Home),
	}
}
