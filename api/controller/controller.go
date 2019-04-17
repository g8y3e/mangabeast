package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController    = newHome()
	mangaController   = newManga()
	chapterController = newChapter()

	errorController = newError()
)

type IController interface {
	setTemplate(name string, template *template.Template)
	registerRoutes()
}

func newController() Controller {
	return Controller{
		template: make(map[string]*template.Template),
	}
}

type Controller struct {
	template map[string]*template.Template
}

func (c *Controller) setTemplate(name string, template *template.Template) {
	c.template[name] = template
}

func (c *Controller) templateByName(name string) *template.Template {
	templateData, ok := c.template[name]
	if !ok {
		panic("Not initiated template: '" + name + "'")
	}
	return templateData
}

func initController(object IController, templateName string, templates map[string]*template.Template) {
	object.setTemplate(templateName, templates[templateName+".html"])
	object.registerRoutes()
}

func Init(templates map[string]*template.Template) {
	initController(homeController, "home", templates)
	initController(mangaController, "manga", templates)
	initController(chapterController, "chapter", templates)

	initController(errorController, "error", templates)

	// public routes
	http.Handle("/image/", http.FileServer(http.Dir("public")))
	http.Handle("/font/", http.FileServer(http.Dir("public")))
	http.Handle("/icon/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
