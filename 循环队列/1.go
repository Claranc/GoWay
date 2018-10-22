package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := []int{4:0}
	var tail int = 0
	var head int = 0
	fmt.Printf("Please input: \n")
	for {
		fmt.Printf(">> ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if len(input.Text()) == 0 {
			fmt.Println("Input error")
			os.Exit(-1)
		}
		par := strings.Split(input.Text(), " ")
		if par[0] == "push" {
			if len(par) < 2 {
				fmt.Println("input error")
				continue
			}
			temp,ok := strconv.Atoi(par[1])
			if ok != nil {
				fmt.Println("string->int failed")
				continue
			}
			flag := Push(&head, &tail, &temp, &a)
			if flag == false {
				fmt.Println("push error")
			} else {
				fmt.Println("OK")
			}
		} else if par[0] == "pop" {
			var temp int
			flag := Pop(&head, &tail, &temp, &a)
			if flag == true {
				fmt.Println(temp)
			} else {
				fmt.Println("pop error")
			}
		} else if par[0] == "help" {
			fmt.Printf("1. push val\n2. pop\n")
		} else {
			fmt.Println("input error")
			continue
		}
	}



}

func Push(head,tail,val *int, a *[]int) bool {
	if(*tail+1)%len(*a) == *head {
		return false
	}
	(*a)[*tail] = *val
	*tail = (*tail+1)%len(*a)
	return true
}

func Pop(head,tail,val *int, a *[]int) bool {
	if *head == *tail {
		return false
	}
	*val = (*a)[*head]
	*head = (*head+1)%len(*a)
	return true
}