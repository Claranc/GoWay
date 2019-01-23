package main

import (
	"log"
	"sync"
)

func main() {
	var lock sync.Mutex
	var mailbox = 0
	var sendcond = sync.NewCond(&lock)
	var recvcond = sync.NewCond(&lock)
	ch := make(chan struct{}, 3)
	max := 6
	go func(max int) {
		defer func() {
			ch <- struct{}{}
		}()
		var count int
		for count < max {
			lock.Lock()
			for mailbox == 1 {
				sendcond.Wait()
			}
			log.Println("mailbox can send mail")
			mailbox = 1
			lock.Unlock()
			recvcond.Broadcast()
			count++
		}
	}(max)

	for i := 0; i < 2; i++ {
		go func(max int) {
			defer func() {
				ch <- struct{}{}
			}()
			var count int
			for count < max {
				lock.Lock()
				for mailbox == 0 {
					recvcond.Wait()
				}
				log.Println("mailbox can read mail")
				mailbox = 0
				lock.Unlock()
				sendcond.Signal()
				count++
			}
		}(max/2)
	}

	<-ch
	<-ch
	<-ch
}
