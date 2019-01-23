package main

import (
	LIST "container/list"
	"fmt"
)

func main() {
	//list
	p := LIST.New()
	p.PushBack(5)
	p.PushBack(6)
	p.PushBack(7)
	for i := p.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
