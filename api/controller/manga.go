package controller

type manga struct {
	Controller
}

func newManga() *manga {
	return &manga{
		newController(),
	}
}

func (m *manga) registerRoutes() {
}
