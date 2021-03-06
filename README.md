# Duck

Duck is a simple but powerful package for learning how to write modular web applications/services in Golang.

Notice: Duck is still in the design stage. It's not working yet. Stay tuned.

# Getting Started

After installing Go and setting up your [GOPATH](http://golang.org/doc/code.html#GOPATH), and then

```go
go get github.com/zddhub/duck
```

# How to use

You can use Duck like this:

```go
package main

import "github.com/zddhub/duck"

func main() {
  d := duck.Incubate()

  d.Get("/", func() string {
    return "Hello world!"
  })

  d.Run()
}
```

Then run your server:
You will now have a Duck webserver running on `localhost:3030`

You can run more complex case:

```go
go run main/main.go
```

# Getting Help

You can read the follow blog to learn more:
* [Go web framework 1](http://www.zddhub.com/memo/2015/07/04/go-web-framework/)
* [Dependency inject](http://www.zddhub.com/memo/2015/07/05/go-dependency-inject/)
* [Go web framwwork 2](http://www.zddhub.com/memo/2015/07/11/go-web-framework-2/)

# About

Inspired by [Martini](https://github.com/go-martini/martini)
