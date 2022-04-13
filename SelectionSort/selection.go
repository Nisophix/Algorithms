package main

import (
	"fmt"
)

func main(){
	arr := []int{64, 25, 12, 22, 11, 14, 1, 3, 2}
	n := len(arr)

	fmt.Println("Unsorted:",arr)

    for i := 0; i < n; i++ {
        var min = i
        for j := i; j < n; j++ {
            if arr[j] < arr[min] {
                min = j
            }
        }
        arr[i], arr[min] = arr[min], arr[i]
    }
    fmt.Println("Sorted:",arr)
}

/*
selectionSort(array, size)
  repeat (size - 1) times
  set the first unsorted element as the minimum
  for each of the unsorted elements
    if element < currentMinimum
      set element as new minimum
  swap minimum with first unsorted position
end selectionSort
*/