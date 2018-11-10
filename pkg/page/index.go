package page

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	Home struct {
		ViewPath string
	}

	HomeService interface {
		Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params)
	}
)

func (h *Home) Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		view := h.ViewPath + "index.html"
		http.ServeFile(w, r, view)
	}
}
