package main

import (
	fl "flag" //包的别名，防止路径不一样名字一样的包起冲突
	fm "fmt"
)


//必须先在这里定义
var p int64
var q int
var str string

func init() {
	//输入Int64类型
	fl.Int64Var(&p, "age", 18, "you should input a age")
	//输入Int类型
	fl.IntVar(&q, "score", 100, "you should input a age")
	//输入string类型
	fl.StringVar(&str, "name", "lmh", "you should input a name")
}

func main() {
	fl.Parse()
	fm.Printf("age: %d\n", p)
	fm.Printf("score: %d\n", q)
	fm.Printf("name: %s\n", str)
}