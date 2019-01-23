package main

import (
	"fmt"
	"errors"
)

func main() {
	a := 0
	b := 5
	res, err := Test(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func Test(a,b int) (c int, err error) {
	if a == 0 {
		err = errors.New("a cannot be 0")
	} else {
		c = b/a
	}
	return
}
