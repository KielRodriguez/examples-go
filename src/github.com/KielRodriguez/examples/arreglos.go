package main

import(
  "fmt"
)

func main() {
  arreleglo := [3]int{1,2,3}

  for i:=0; i < len(arreleglo); i++ {
    fmt.Println(arreleglo[i])
  }
}
