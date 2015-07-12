package main

import (
	"net/http"
)

type Router interface {
	Get(string, ...Handler) Router
	Post(string, ...Handler) Router
	Handle(http.ResponseWriter, *http.Request, *Context)
}

// route entry
type route struct {
	method   string
	pattern  string
	handlers []Handler
}

type router struct {
	routes []*route
}

func NewRouter() Router {
	r := make([]*route, 0)
	return &router{r}
}

func (r *router) Handle(w http.ResponseWriter, req *http.Request, c *Context) {
	rc := &RouterContext{Injector: c.Injector, handlers: r.routes[0].handlers, index: 0}
	rc.Run()
}

func (r *router) Get(pattern string, handlers ...Handler) Router {
	return r.addRoute("GET", pattern, handlers)
}

func (r *router) Post(pattern string, handlers ...Handler) Router {
	return r.addRoute("POST", pattern, handlers)
}

func (r *router) addRoute(method string, pattern string, h []Handler) Router {
	handlers := make([]Handler, 0)
	handlers = append(handlers, h...)
	rt := &route{method, pattern, handlers}
	r.routes = append(r.routes, rt)
	return r
}
