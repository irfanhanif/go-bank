package setup

import (
	"github.com/irfanhanif/go-bank-interface/pkg/http/web"
)

// SetupWebObserver function is a function to setup the Web Observer to map between Path request and the handler methods.
func SetupWebObserver() *web.WebObserver {
	webObserver := web.BuildWebObserver(
		web.WithRouter(),
	)
	return webObserver
}
