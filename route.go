package main

import (
	"fmt"
	"net/http"
	"regexp"
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
	regex    *regexp.Regexp
}

type router struct {
	routes []*route
}

func NewRouter() Router {
	r := make([]*route, 0)
	return &router{r}
}

func (r *router) Handle(w http.ResponseWriter, req *http.Request, c *Context) {
	rt, params := r.MatchRoute(req)
	if rt == nil {
		return
	}
	c.SetMap(params)
	rc := &RouterContext{Injector: c.Injector, handlers: rt.handlers, index: 0}
	rc.Run()
}

type Params map[string]string

// Return match route
func (r *router) MatchRoute(req *http.Request) (*route, Params) {
	for key, val := range r.routes {
		if val.method == req.Method {
			if match, params := val.Match(req.URL.Path); match {
				return r.routes[key], params
			}
		}
	}
	return nil, nil
}

func (r route) Match(path string) (bool, map[string]string) {
	matches := r.regex.FindStringSubmatch(path)
	if len(matches) > 0 && matches[0] == path {
		params := make(map[string]string)
		for i, name := range r.regex.SubexpNames() {
			if len(name) > 0 {
				params[name] = matches[i]
			}
		}
		return true, params
	}
	return false, nil
}

var routeReg1 = regexp.MustCompile(`:[^/#?()\.\\]+`)
var routeReg2 = regexp.MustCompile(`\*\*`)

func newRoute(method string, pattern string, handlers []Handler) *route {
	route := route{method, pattern, handlers, nil}
	pattern = routeReg1.ReplaceAllStringFunc(pattern, func(m string) string {
		return fmt.Sprintf(`(?P<%s>[^/#?]+)`, m[1:])
	})
	var index int
	pattern = routeReg2.ReplaceAllStringFunc(pattern, func(m string) string {
		index++
		return fmt.Sprintf(`(?P<_%d>[^#?]*)`, index)
	})
	pattern += `\/?`
	route.regex = regexp.MustCompile(pattern)
	return &route
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
	rt := newRoute(method, pattern, handlers)
	r.routes = append(r.routes, rt)
	return r
}
