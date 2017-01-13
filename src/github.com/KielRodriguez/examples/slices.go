package main

import(
  "fmt"
)

/*
Slice tiene como base un arreglo
estructura basada en los arreglos
pero a diferencia tiene mas flexibilidad
esto quiere decir que se uede redimencionar
*/
func main() {
    //al no inicializar sin tamaño en arreglos
    // esto pasa a ser un slicer
    matriz := []int{1,2,3,4}

    if matriz == nil {
      fmt.Println("matriz esta vacio")
    } else {
        //slicing parte un arreglo de acuerdo a los
        //parametros indicados
        //[1:2] imprime solo 2
        //[:2]  del inicio a la pasicion 2
        //[2:] de la posicion 2 al final
        fmt.Println( matriz[1:2] ) 
    }


}
