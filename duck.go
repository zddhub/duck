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
	handlers []Handler // handler all fanc
	logger   *log.Logger
	IP       string
	Port     string
}

func Incubate() *Duck {
	return &Duck{Injector: New(), IP: "", Port: "3030",
		logger: log.New(os.Stdout, "\033[0;32;34m[duck] \033[m", 0)}
}

func (d *Duck) Run() {
	d.logger.Println("listening on", d.IP+":"+d.Port)
	d.logger.Fatalln(http.ListenAndServe(d.IP+":"+d.Port, d))
}

func (d *Duck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[Duck] start ServeHTTP")
	d.SetMap(w)
	d.SetMap(r)

	for i := 0; i < len(d.handlers); i++ {

		if ret, err := d.Invoke(d.handlers[i]); err == nil {
			if len(ret) != 0 {
				fmt.Fprintf(w, ret[0].Interface().(string))
			}
		} else {
			fmt.Println(err)
		}
	}

	fmt.Println("[Duck] end   ServeHTTP")
}

func (d *Duck) Get(pattern string, handler Handler) {
	fmt.Println("[Duck] Get")
	d.handlers = append(d.handlers, handler)
}
