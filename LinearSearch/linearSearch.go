package main

import (
	"fmt"
)

func main(){
	//the array doesn't have to be sorted
	arr := []int{775,876,984,1002,1003,204,30,3,22,133,242,344,545,666,667};
	result := linearSearch(arr, len(arr), 667)
	if result != -1{
        fmt.Println(result)
    }

}

func linearSearch(arr []int, n, x int) int {
	for i := 0; i<n; i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}