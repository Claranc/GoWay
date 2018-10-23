package main

import "fmt"

type Student struct {
	age int
	name string
}

type Teacher struct {
	base *Student
	course string
}

func (a Student) GetInfo() {
	fmt.Printf("name:%s, age:%d\n", a.name, a.age)
}

func (a Teacher) GetInfo() {
	fmt.Printf("name:%s, age:%d, course:%s\n", a.base.name, a.base.age, a.course)
}

func (a Teacher) GetName() {
	fmt.Printf("name:%s\n", a.base.name)
}

//将结构体共同的函数放进接口里。共同的函数是指有相同的名称、参数列表（不包括参数名）以及返回值。
type Man interface {
	GetInfo()
}

func main() {
	var t Man
	stu1 := Student{23, "fxj"}
	t = stu1
	t.GetInfo()
	te1 := Teacher{&Student{35, "glj"}, "english"}
	t = te1
	t.GetInfo()
}
