package main

import(
  "fmt"
)

func main() {
  slice := make([]int, 3, 6)
  slice = append(slice, 2)
  slice = append(slice, 3)
  slice = append(slice, 4)
  slice = append(slice, 5)
  slice = append(slice, 6)
  fmt.Println(slice)
}
