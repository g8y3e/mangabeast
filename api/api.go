package api

import (
	"github.com/g8y3e/mangabeast/api/controller"
	"log"
	"net/http"
)

var (
	port = ":3000"
)

func Init() {
	templates := populateTemplates()
	controller.Init(templates)

	log.Println("Init server on port:", port)
	http.ListenAndServe(port, nil)
}
