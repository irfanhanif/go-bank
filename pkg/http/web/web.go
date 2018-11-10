package web

import (
	"errors"
	"github.com/irfanhanif/go-bank-interface/pkg/page"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	Option func(*WebObserver) error

	WebObserver struct {
		Router *httprouter.Router
	}
)

const (
	// viewPath is variable to locate the view directory.
	viewPath = "../../web/view/"
)

var (
	// Put page list here before register the method into mapping.
	home  = &page.Home{viewPath + "home/"}
	login = &page.Login{viewPath + "login/"}

	// Map of http path to the handler function. Method GET.
	httpGetHandler = map[string]func(w http.ResponseWriter, t *http.Request, _ httprouter.Params){
		"/":      home.Index(),
		"/login": login.Index(),
	}

	// Map of http path to the handler function. Method POST.
	httpPostHandler = map[string]func(w http.ResponseWriter, t *http.Request, _ httprouter.Params){

	}
)

// Web observer builder. Build functions with Option type return.
func BuildWebObserver(opts ...Option) *WebObserver {
	obsvr := new(WebObserver)
	for _, fn := range opts {
		if err := fn(obsvr); err != nil {
			panic(err)
		}
	}
	return obsvr
}

// WithRouter is a function to register all the routings.
func WithRouter() Option {
	return func(obsvr *WebObserver) error {
		if obsvr == nil {
			return errors.New("Observer is not set")
		}

		router := httprouter.New()
		for path, method := range httpGetHandler {
			router.GET(path, method)
		}
		for path, method := range httpPostHandler {
			router.POST(path, method)
		}

		obsvr.Router = router

		return nil
	}
}
