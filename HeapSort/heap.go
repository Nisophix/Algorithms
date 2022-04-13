package main

import(
    "fmt"
)

func main(){
    arr := []int{64, 25, 12, 22, 11, 14, 1, 3, 2}
    n := len(arr)
    fmt.Println("Unsorted:",arr)
    sort(arr, n)
    fmt.Println("Sorted:",arr)
}

func heapify(arr []int, n, i int){
    largest := i
    l := 2 * i + 1
    r := 2 * i + 2
    if l < n && arr[l] > arr[largest]{
        largest = l
    }
    if r < n && arr[r] > arr[largest]{
        largest = r
    }
    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

func sort(arr []int, n int){
    for i := n / 2 - 1; i >= 0; i-- {
        heapify(arr, n, i)
    }
    for i := n - 1; i > 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }
}