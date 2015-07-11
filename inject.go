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
	val := inj.mappers[t]
	if val.IsValid() {
		return val
	}

	// if t is Interface, find some implements one.
	if t.Kind() == reflect.Interface {
		for k, v := range inj.mappers {
			if k.Implements(t) {
				val = v
				break
			}
		}
	}

	return val
}

func (inj *Injector) Invoke(i interface{}) ([]reflect.Value, error) {
	t := reflect.TypeOf(i)
	if t.Kind() != reflect.Func {
		panic("Should invoke a function!")
	}
	inValues := make([]reflect.Value, t.NumIn())
	for k := 0; k < t.NumIn(); k++ {
		val := inj.Get(t.In(k))
		if !val.IsValid() {
			return nil, fmt.Errorf("Value not found for type %v", t.In(k))
		}
		inValues[k] = val
	}
	ret := reflect.ValueOf(i).Call(inValues)
	return ret, nil
}
