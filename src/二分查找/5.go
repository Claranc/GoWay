package main

import "fmt"

func main() {
	ee := []int {1,2,3,4,5}
	f := ee
	f[1]= 89
	fmt.Println(ee)
	a := string("12345")
	var b =  []byte(a)
	fmt.Printf("%c", b)
	b[2] = '0'
	fmt.Printf("%c",b)
}
