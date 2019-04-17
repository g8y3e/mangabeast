package controller

import (
	"github.com/g8y3e/mangabeast/api/model"
	"github.com/g8y3e/mangabeast/api/viewmodel"
	"net/http"
)

type home struct {
	Controller
}

func newHome() *home {
	return &home{
		newController(),
	}
}

func (h *home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
}

func (h *home) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	if r.URL.Path != "/" && r.URL.Path != "/home" {
		errorController.handleError("404", w, r)
		return
	}

	homeData, err := model.GetHome()
	if err != nil {
		errorController.handleError("404", w, r)
		return
	}

	vm := viewmodel.NewHome(homeData)
	h.templateByName("home").Execute(w, vm)
}
