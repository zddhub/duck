package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "reflect"
)

type Handler interface{}

type Duck struct {
	*Injector
	handlers      []Handler // handler all fanc
	routerHandler Handler   // handle a route enter <-
	logger        *log.Logger
	IP            string
	Port          string
}

type MatureDuck struct {
	*Duck
	Router
}

func Incubate() *MatureDuck {
	d := &Duck{Injector: New(), IP: "", Port: "3030",
		logger: log.New(os.Stdout, "\033[0;32;34m[duck] \033[m", 0)}
	r := NewRouter()
	d.routerHandler = r.Handle
	return &MatureDuck{d, r}
}

func (d *Duck) Run() {
	d.logger.Println("listening on", d.IP+":"+d.Port)
	d.logger.Fatalln(http.ListenAndServe(d.IP+":"+d.Port, d))
}

func (d *Duck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Duck] start ServeHTTP")
	d.createContext(w, r).Run()
	fmt.Println("[Duck] end   ServeHTTP")
}

func (d *Duck) createContext(w http.ResponseWriter, r *http.Request) *Context {
	d.SetMap(w)
	d.SetMap(r)
	c := &Context{Injector: d.Injector, routerHandler: d.routerHandler, handlers: d.handlers, index: 0}
	d.SetMap(c)
	return c
}
