package main

import "fmt"
type Persona struct {
	nombre string 
	estatura float64
}

func (p *Persona) correr() string{
	return p.nombre+" esta corriendo"
}

func main() {	
	p:=Persona{"Jhon", 1.72}
	fmt.Println(p.correr())
}

