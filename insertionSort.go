package main

import "fmt"

func insertionSort(slice []int) {
  for i, v := range slice {
    if (i == 0) { continue }
    key := v
    j := i-1
    for {
      if (j < 0 || slice[j] < key) { break }
      slice[j+1] = slice[j]
      j--
    }
    slice[j+1] = key
  }
}

func main() {
  slice := []int{5, 4, 3, 2, 1}
  insertionSort(slice)
  fmt.Println(slice)
}
