package page

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	Login struct {
		ViewPath string		// location of view of this page
	}

	// This interface is needed for Mocking purpose.
	LoginService interface {
		Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params)
	}
)

// Index function send a default index.html to the client side.
func (l *Login) Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		view := l.ViewPath + "index.html"
		http.ServeFile(w, r, view)
	}
}
