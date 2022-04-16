package main

import (
    "fmt"
)

func main(){
    //the array has to be sorted if it is not already
    arr := []int{0,3,22,133,242,344,545,666,667,775,876,984,1002,1003,2043};
    result := binarySearch(arr, 0, len(arr), 775);
    if result != -1{
        fmt.Println(result)
    }
}

func binarySearch(arr []int, from, to, x int) int {
    for from <= to {
        mid := from + (to - from) / 2;
        fmt.Println("mid: ",mid)
        if arr[mid] == x{
            return mid;
        }else if arr[mid] < x {
            from = mid + 1;
        }else {
            to = mid - 1;
        }
    }
    return -1;
}
 