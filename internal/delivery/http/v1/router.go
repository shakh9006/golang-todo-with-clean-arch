package router

type Router interface {
	InitAndServeRoutes(string) error
}
