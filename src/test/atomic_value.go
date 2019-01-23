package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var box atomic.Value
	box2 := box // 原子值在真正使用前可以被复制。
	v1 := [...]int{1, 2, 3}
	v2 := "123"
	box.Store(v1)
	box2.Store(v2)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	fmt.Println()

	box3 := box // 原子值在真正使用后不应该被复制！
	fmt.Printf("The value load from box3 is %v.\n", box3.Load())
	//box3.Store(v3) // 这里会引发一个panic，报告存储值的类型不一致。
	_ = box3 //没使用box3
	fmt.Println()

	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	box6.Store(v6)
	v6[1] = 4 // 注意，此处的操作会改变box6
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
	// 正确的做法如下，复制一个store进去
	v6 = []int{1, 2, 3}
	store := func(v []int) {
		replica := make([]int, len(v))
		copy(replica, v)
		box6.Store(replica)
	}
	store(v6)
	v6[2] = 5 // 此处的操作是安全的。
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
}

