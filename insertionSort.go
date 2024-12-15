package main

import "fmt"

func insertionSort(slice []int) {
  for i := 0; i < len(slice); i++ {
    for j := i; j > 0 && slice[j-1] > slice[j]; j-- {
      slice[j], slice[j-1] = slice[j-1], slice[j]
    }
  }
}

func main() {
  slice := []int{5, 4, 3, 2, 1}
  insertionSort(slice)
  fmt.Println(slice)
}
