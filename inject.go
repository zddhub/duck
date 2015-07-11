package main

import (
	"fmt"
	"reflect"
)

type Injector struct {
	mappers map[reflect.Type]reflect.Value // 根据类型map实际的值
}

func New() *Injector {
	return &Injector{make(map[reflect.Type]reflect.Value)}
}

func (inj *Injector) SetMap(value interface{}) {
	inj.mappers[reflect.TypeOf(value)] = reflect.ValueOf(value)
}

func (inj *Injector) Get(t reflect.Type) reflect.Value {
	return inj.mappers[t]
}

func (inj *Injector) Invoke(i interface{}) []reflect.Value {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Func {
		panic("Should invoke a function!")
	}
	inValues := make([]reflect.Value, t.NumIn())
	for k := 0; k < t.NumIn(); k++ {
		inValues[k] = inj.Get(t.In(k))
	}
	ret := reflect.ValueOf(i).Call(inValues)
	return ret
}
