package app

import (
	"github.com/r0fls/reinhardt/src/url"
	"github.com/r0fls/reinhardt/test/app/views"
)

func Urls() url.Urls {
	return []url.Url{
		url.URL("/bar", views.Home),
		url.URL("/$", views.Home),
	}
}
