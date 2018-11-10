package page

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	Login struct {
		ViewPath string
	}

	LoginService interface {
		Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params)
	}
)

func (l *Login) Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		view := l.ViewPath + "index.html"
		http.ServeFile(w, r, view)
	}
}