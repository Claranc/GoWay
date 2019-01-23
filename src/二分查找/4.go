package main

import "fmt"

func main() {
	a := []int{1,2,3,4}
	FullRank(&a, 4)
}

func FullRank(a *[]int, k int) {
	if k == 1 {
		fmt.Println((*a))
		return
	}
	for i := 0; i < k; i++ {
		(*a)[i], (*a)[k-1] = (*a)[k-1], (*a)[i]
		FullRank(a, k-1)
		(*a)[i], (*a)[k-1] = (*a)[k-1], (*a)[i]
	}
	return
}
