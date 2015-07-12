package main

import (
	"fmt"
	"net/http"
)

type Context struct {
	*Injector
	handlers      []Handler
	routerHandler Handler
	index         int
}

func (c *Context) Next() {
	c.index += 1
	c.Run()
}

func (c *Context) Run() {
	// Run all Duck handles
	for c.index < len(c.handlers) {

		if ret, err := c.Invoke(c.handlers[c.index]); err == nil {
			if len(ret) != 0 {
				rv := c.Get(c.GetType((*http.ResponseWriter)(nil)))
				res := rv.Interface().(http.ResponseWriter)
				fmt.Fprintf(res, ret[0].Interface().(string))
			}
		} else {
			fmt.Println(err)
		}
		c.index++
	}
	// Run all route handles
	if c.index == len(c.handlers) {
		c.Invoke(c.routerHandler)
	}
}

type RouterContext struct {
	*Injector
	handlers []Handler // all route handler
	index    int
}

func (rc *RouterContext) Run() {
	for i := 0; i < len(rc.handlers); i++ {
		if ret, err := rc.Invoke(rc.handlers[i]); err == nil {
			if len(ret) != 0 {
				rv := rc.Get(rc.GetType((*http.ResponseWriter)(nil)))
				res := rv.Interface().(http.ResponseWriter)
				fmt.Fprintf(res, ret[0].Interface().(string))
			}
		} else {
			fmt.Println(err)
		}
	}
}
