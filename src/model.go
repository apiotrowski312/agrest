package agrest

type agrest struct {
	baseURL string
	timeout int
}

func CreateClient() agrest {
	return agrest{}
}

func (a *agrest) SetBaseURL(baseURL string) {
	a.baseURL = baseURL
}
