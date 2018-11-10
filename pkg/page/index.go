package page

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type (
	Home struct {
		ViewPath string		// location of view of this page
	}

	// This interface is needed for Mocking purpose.
	HomeService interface {
		Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params)
	}
)

// Index function send a default index.html to the client side.
func (h *Home) Index() func(w http.ResponseWriter, t *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		view := h.ViewPath + "index.html"
		http.ServeFile(w, r, view)
	}
}
