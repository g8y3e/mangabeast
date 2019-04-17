package viewmodel

type HeaderInfo struct {
}

type FooterInfo struct {
}

type ViewModel struct {
	Active string

	Header HeaderInfo
	Footer FooterInfo
}

func newViewModel(active string) ViewModel {
	result := ViewModel{
		Active: active,
	}
	/*layout, err := model.GetLayout()
	if err != nil {
		return result
	}
	*/
	return result
}
