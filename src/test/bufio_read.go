package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	//读
	var s = new(bytes.Buffer) //指针方法
	var src = []byte{'d','r','t','o'}
	s.Write(src)
	a := bufio.NewReader(s) //默认4096字节，4KB
	//a := bufio.NewReaderSize(s,2) //手动分配空间
	var p = make([]byte, 3)
	a.Read(p)
	for _,v := range p {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	//写
	b := bufio.NewWriterSize(s,2)
	//写入内容塞满了缓冲区会自动刷新到底层
	b.WriteString("qwerdfb")
	fmt.Println(s.String())
	b.Flush() //把最后剩余在缓冲区的内容刷新到底层
	fmt.Println(s.String())
}
