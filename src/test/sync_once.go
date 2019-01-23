package main

import (
	"fmt"
	"sync"
)

func main() {
	var test sync.Once
	var wg sync.WaitGroup
	wg.Add(10)
	p := func () {
		fmt.Println("done")
	}
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			test.Do(p)
		}()
	}
	wg.Wait()
}


