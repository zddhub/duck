package main

import "fmt"

func main() {
	d := Incubate()
	d.Get("/", func() string {
		fmt.Println("Hello world!")
		return "Helle world!"
	})
	d.Run()
}
