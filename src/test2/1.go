package main

import "fmt"

func main() {
	p := bibao(5)
	fmt.Println(p(1)) //输出7
	fmt.Println(p(1)) //输出8
}

func bibao(i int) func(i int) int {
	var x int
	x = i //自由变量x
	a := func(j int) int {
		x++
		return x + j
	}
	return a //返回闭包函数
}