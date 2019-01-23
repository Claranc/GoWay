package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString("chuanboniubi")
	p := make([]byte, 5)
	buf.Read(p)
	for _,v := range p {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	buf.Write([]byte("julao"))
	q := buf.Next(5)
	for _,v := range q {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	c,_ := buf.ReadByte()
	fmt.Printf("%c", c)
	fmt.Println()

	fmt.Println(buf.String())
	fmt.Println(string(buf.Next(10)))
}
