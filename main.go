package main

func main() {
	d := Incubate()
	d.Get("/", func() string {
		return "Helle world!"
	})
	d.Run()
}
