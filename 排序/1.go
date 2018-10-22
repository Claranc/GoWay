package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{5,8,9,3,2,15,78,0,63,1}
	b := QuickSort(a)
	fmt.Println(b)
}

//冒泡排序
func Maopao(a []int) []int {
	for j := 0; j < len(a); j++ {
		flag := false
		for i := 0; i < len(a)-1-j; i++ {
			if a[i] < a[i+1] {
				flag = true
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		if flag == false {
			break
		}
	}
	return a
}

//插入排序
func Charu(a []int) []int {
	for i := 1; i < len(a); i++ {
		val := a[i]
		j := i-1
		for ; j >= 0; j-- {
			if a[j] > val {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = val
	}
	return a
}

//选择排序
func Xuanze(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := a[i]
		var location int = i
		for j := i; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				location = j
			}
		}
		a[i],a[location] = a[location], a[i]
	}
	return a
}

//归并排序
func Guibing(a []int) []int {
	SplitSort(a, 0, len(a)-1)
	return a
}

func SplitSort(a []int, p,r int) {
	if p <r {
		q := (p + r)/2
		SplitSort(a, p, q)
		SplitSort(a, q+1, r)
		Merge(a, p, q, r)
	}
}

func Merge(a []int, p,q,r int) {
	n1 := q+1-p
	n2 := r-q
	temp1 := make([]int, n1+1)
	temp2 := make([]int, n2+1)
	for i := 0; i < n1; i++ {
		temp1[i] = a[p+i]
	}
	temp1[n1] = math.MaxInt64
	for i := 0; i < n2; i++ {
		temp2[i] = a[q+1+i]
	}
	temp2[n2] = math.MaxInt64
	i := 0
	j := 0
	for index := p; index < r+1; index++{
		if temp1[i] < temp2[j] {
			a[index] = temp1[i]
			i++
		}else {
			a[index] = temp2[j]
			j++
		}
	}
}

//快速排序
func QuickSort(a []int) []int {
	Split(a, 0, len(a)-1)
	return a
}

func Split(a []int, p,r int) {
	if p < r {
		q := Partion(a, p, r)
		Split(a, p, q-1)
		Split(a, q+1, r)
	}
}

func Partion(a []int, p,r int) int {
	val := a[r]
	count := p
	for i := p; i <= r-1; i++ {
		if a[i] < val {
			a[i], a[count] = a[count], a[i]
			count++
		}
	}
	a[count], a[r] = a[r], a[count]
	return count
}
