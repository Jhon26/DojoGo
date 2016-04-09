package main

import "fmt"

func main() {
  fmt.Print("Ingrese el texto: ")
  var entrada string
  fmt.Scanf("%s", &entrada)
  fmt.Println(entrada)
}