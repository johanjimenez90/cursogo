package main

import "fmt"

func main() {
	nombre, apellido := "Alexys", "Lozada"
	segundo_nombre, segundo_apellido := "johan", "jimenez"
	nombre, edad := "Pedro", 20
	segundo_nombre, segunda_edad := "jessica", 19

	fmt.Println(nombre, apellido, edad)
	fmt.Println(segundo_nombre, segundo_apellido, segunda_edad)
}
