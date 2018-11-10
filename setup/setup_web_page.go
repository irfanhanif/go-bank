package setup

import (
	"github.com/irfanhanif/go-bank-interface/pkg/http/web"
)

func SetupWebObserver() *web.WebObserver {
	webObserver := web.BuildWebObserver(
		web.WithRouter(),
	)
	return webObserver
}