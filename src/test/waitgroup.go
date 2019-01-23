package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer func() {
				wg.Done()
			}()
			fmt.Println("666")
		}()
	}
	wg.Wait()
}
