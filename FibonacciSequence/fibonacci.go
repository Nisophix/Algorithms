package main

import (
	"fmt"
)

func main(){
	printFib(20)
}

func printFib(n int){
	var a, b, c int = 0, 1, 0
	fmt.Printf("%d ", a)
	for i := 0; i < n-1; i++ {
		c = a+b
		a = b
		b = c
		fmt.Printf("%d ", a)
	}
	fmt.Printf("\n")
}