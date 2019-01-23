package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 20)
	go DP(ch, 10)
	time.Sleep(1*time.Second)
	/*for {
		v,ok :=<-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}*/
	for v := range ch {
		fmt.Println(v)
	}
}


func DP(ch chan int, i int) {
	for j := 0; j < i; j++ {
		ch <- j*10
	}
	close(ch)
}