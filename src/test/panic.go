package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

func main() {
	a := 0
	b := 5
	defer func() {
		if c := recover(); c != nil{
			log.Println(c)
		}
	}()
	if a == 0 {
		panic(errors.New("a cannot be 0"))
	}
	fmt.Println(b/a)
}
