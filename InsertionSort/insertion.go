package main

import (
    "fmt"
)

func main(){
    arr := []int{64, 25, 12, 22, 11, 14, 1, 3, 2}
    n := len(arr)

    fmt.Println("Unsorted:",arr)

    for i := 1; i < n; i++{
        key := arr[i]
        j := i - 1;

        for j >= 0 && arr[j] > key{
            arr[j + 1] = arr[j]
            j = j - 1
        }
        arr[j + 1] = key
    }
    fmt.Println("Sorted:",arr)
}