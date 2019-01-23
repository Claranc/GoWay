package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i:=0;i < 1000; i++{
		go Print(i, ch)
	}
	for {
		num := <- ch
		fmt.Printf("I come from %d\n", num)
	}
	//time.Sleep(1000*time.Microsecond) //主程序退出了其他的也停止了
}

func Print(k int, ch chan int) {
	for{
		ch <- k
	}
}