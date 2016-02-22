package app

import (
	"github.com/r0fls/reinhardt/src/url"
	"github.com/r0fls/reinhardt/test/app/views"
)

func Urls() []url.Url {
	return []url.Url{
		url.Url{"/bar", views.Home},
		url.Url{"/", views.Home},
	}
}
