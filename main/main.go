package main

import (
	"fmt"
	"net/http"
  "github.com/zddhub/duck"
)

func main() {
	d := duck.Incubate()

	d.Get("/", func() string {
		return "Hello world!"
	})

	d.Get("/zdd", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Duck!")
	})

	d.Get("/zddhub/:id", func(params duck.Params) string {
		return "Hello " + params["id"] + " "
	}, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zddhub - 1 ")
	}, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zddhub - 2 ")
	}, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "zddhub - 3 ")
	})

	d.Run()
}
