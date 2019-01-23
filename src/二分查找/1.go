package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//在不重复的数组中找出val
func main() {
	rand.Seed(time.Now().Unix())
	p := []int{}
	for i := 0; i < 5; i++ {
		p = append(p, rand.Intn(10))
	}
	sort.Ints(p)
	fmt.Println(p)
	flag := BinarySearch(3, p)
	fmt.Println(flag)

}

func BinarySearch(k int, p []int) bool {
	low, high := 0, len(p)-1
	for low <= high {
		mid := low+ (high-low)/2  //用这个式子可以防止溢出
		if p[mid] == k {
			return true
		}
		if k < p[mid] {
			high = mid -1
		} else {
			low = mid + 1
		}
	}
	return false
}