package main

import "fmt"

//找出数组中第一个大于等于val
func main() {
	a := []int{1,2,2,2,4,5,5,5,8,9,9,9,10}
	location := FindFirstBigger(a, 3)
	fmt.Println(location)

}

func FindFirstBigger(a []int, val int) int {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := low + (high - low) / 2
		if a[mid] < val {
			low = mid + 1
		}
		if a[mid] >= val {
			if mid == 0 {
				return mid
			}
			if a[mid - 1] >= val {
				high = mid - 1
			} else {
				return mid
			}
		}
	}
	return -1
}