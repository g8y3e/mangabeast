package controller

type chapter struct {
	Controller
}

func newChapter() *chapter {
	return &chapter{
		newController(),
	}
}

func (ch *chapter) registerRoutes() {
}
