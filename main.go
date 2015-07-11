package main

import (
	"fmt"
	"net/http"
)

func main() {
	d := Incubate()

	d.Get("/", func() string {
		fmt.Println("Hello world!")
		return "Hello world!"
	})

	d.Get("/zdd", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Duck!")
	})

	d.Run()
}
