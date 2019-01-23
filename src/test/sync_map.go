package main

import (
	"fmt"
	"reflect"
	"sync"
)

type NameMap struct {
	mp sync.Map
	KeyType reflect.Type
	ValueType reflect.Type
}

func (this *NameMap)Store(key interface{}, value interface{}) {
	if reflect.TypeOf(key) != this.KeyType || reflect.TypeOf(value) != this.ValueType {
		return
	}
	this.mp.Store(key, value)
}

func (this *NameMap)Load(key interface{}) (interface{}, bool) {
	if reflect.TypeOf(key) != this.KeyType {
		return nil, false
	}
	return this.mp.Load(key)
}

func (this *NameMap)Delete(key interface{}) {
	if reflect.TypeOf(key) != this.KeyType {
		return
	}
	this.mp.Delete(key)
}

func main() {
	var NP = new(NameMap)
	NP.KeyType = reflect.TypeOf(string("fxj"))
	NP.ValueType = reflect.TypeOf(int(123))
	NP.Store("fxj", 95)
	NP.Store("lmh",96)
	v,ok1 := NP.Load("fxj")
	if ok1 {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
	NP.Delete("fxj")
	p,ok2 := NP.Load("fxj")
	if ok2 {
		fmt.Println(p)
	} else {
		fmt.Println("not found")
	}
}
