package main

import (
  "fmt"
)

/*
  igual que         ==
  mayor que         >
  menor que         <
  menor o igual que <=
  meyor o igual que >=
  AND               &&
  OR                ||
*/
func main() {

  var edad int
  fmt.Println("¿Cual es tu edad?")

  fmt.Scanf("%d", &edad)

  if edad > 30  {
    fmt.Println("Tu ya eres un adulto mayor")
  } else if edad >= 18 {
    fmt.Println("Eres mayor de edad")
  } else {
    fmt.Println("Eres un niño")
  }

}
