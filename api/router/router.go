package api

type IRouter interface {
}

type Router struct {
}

func NewRouter() IRouter {
	return &Router{}
}
