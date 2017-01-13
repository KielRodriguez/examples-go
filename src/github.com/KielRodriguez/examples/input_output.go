package main

import(
  "fmt"
  "bufio"
  "os"
)

func main()  {
  // fmt.Print -> imprime sin salto de linea
  // fmt.Println -> imprime con un salto de linea al final
  // fmt.Printf -> imprime con formato (interpolación)

  reader := bufio.NewReader(os.Stdin);
  fmt.Println("Ingresa nombre:")
  nombre, err := reader.ReadString('\n')

  if err == nil {
    fmt.Printf("Su nombre %s", nombre)
  } else {
    fmt.Println(err)
  }
}
