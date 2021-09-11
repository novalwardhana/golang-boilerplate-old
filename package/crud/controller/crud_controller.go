package controller

type controller struct {
}

type Controller interface {
}

func NewController() Controller {
	return &controller{}
}
