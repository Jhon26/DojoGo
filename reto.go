package main

import "fmt"

func main() {
	var entrada string
  	fmt.Scanf("%s", &entrada)
  	s:=len(entrada)
  	for i:=1; i<=s; i++{
 		fmt.Println(entrada[0:i])	
  	}
  	for i:=s; i>0; i--{
 		fmt.Println(entrada[0:i])	
  	}
  	
}
