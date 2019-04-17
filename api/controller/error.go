package controller

import (
	"github.com/g8y3e/mangabeast/api/viewmodel"
	"net/http"
)

type errController struct {
	Controller
}

func newError() *errController {
	return &errController{
		newController(),
	}
}

func (e *errController) registerRoutes() {
}

func (e *errController) handleError(errorName string, w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewError(errorName)
	e.templateByName("error").Execute(w, vm)
}
