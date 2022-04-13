package main

import (
    "fmt"
)

func main(){
    arr := []int{64, 25, 12, 22, 11, 14, 1, 3, 2}
    n := len(arr)-1

    fmt.Println("Unsorted:",arr)

    for i := 0; i < n; i++ {
        for j := 0; j < n-i; j++{
            if arr[j] > arr[j+1]{
                arr[j], arr[j+1] = arr[j+1], arr[j]
            } 
        } 
    }
    fmt.Println("Sorted:",arr)
}

/*
Starting from the first index, compare the first and the second elements.
If the first element is greater than the second element, they are swapped.
Now, compare the second and the third elements. Swap them if they are not in order.
The above process goes on until the last element. 
*/