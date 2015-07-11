package main

import (
	"fmt"
	"net/http"
	"reflect"
)

type Duck struct {
	IP   string
	Port string
}

func Incubate() Duck {
	return Duck{"", "3030"}
}

func (d Duck) Run() {
	fmt.Println("[Duck] listening on", d.IP+":"+d.Port)
	http.ListenAndServe(d.IP+":"+d.Port, nil)
}

//: 1
// func (d Duck) Get(pattern string, f func() string) {
//  http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, f()) // 将f()的返回值写到w
//  })
// }

//: 2
func (d Duck) Get(pattern string, i interface{}) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Func {
		return
	}

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var in = make([]reflect.Value, t.NumIn())
		if t.NumIn() == 2 {
			in[0], in[1] = reflect.ValueOf(w), reflect.ValueOf(r)
		}
		ret := v.Call(in)
		if len(ret) != 0 {
			fmt.Fprintf(w, ret[0].Interface().(string))
		}
	})

}
