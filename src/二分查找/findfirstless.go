package main

import "fmt"

//找出数组中最后一个小于等于val的坐标
func main() {
	a := []int{1,2,2,2,4,5,5,5,8,9,9,9,10}
	location := FindFirstLess(a, 3)
	fmt.Println(location)

}

func FindFirstLess(a []int, val int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if a[mid] <= val {
			if mid == len(a) - 1 || a[mid + 1] > val {
				return mid
			} else {
				low = mid - 1
			}
		}
		if a[mid] > val {
			high = mid - 1
		}
	}
	return -1
}
