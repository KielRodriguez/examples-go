package main

import (
  "fmt"
)

func main () {

  var numero int

  fmt.Println("Especifique un número para obtener los pares: ")
  fmt.Scanf("%d\n", &numero )
  for i:=1; i <= numero; i++ {
    if i % 2 == 0 {
      fmt.Printf("Linea %d\n", i)
    }
  }

  fmt.Println("=========================")
  cont := 0
  for cont < 10
    cont++
    if cont == 2 {
      continue
    } else if cont == 8 {
      break
    }
    fmt.Println(cont)

  }
}
