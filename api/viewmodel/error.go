package viewmodel

type Error struct {
	Title string
	Body  string

	ViewModel
}

func NewError(errName string) Error {
	switch errName {
	case "404":
		return Error{
			Title:     "Not Found",
			Body:      "Page not found error",
			ViewModel: newViewModel("error"),
		}
	}

	return Error{}
}
