package main

import (
	"fmt"
	"strings"
)

func main() {

	var entero int = 10
	var flotante float64 = 3.14
	var booleano bool = true
	var cadena string = "Hola, Mundo"
	var enMayus string = strings.ToUpper(cadena)

	// otra forma de declarar variables
	// entero := 10
	// flotante := 3.14
	// booleano := true
	// cadena := "Hola, Mundo"

	fmt.Println(entero)
	fmt.Println(flotante)
	fmt.Println(booleano)
	fmt.Println(cadena)
	fmt.Println(enMayus)

	// arrays y slices
	array := [3]int{1, 2, 3}
	slice := []int{4, 5, 6}
	increment := append(slice, 7)

	fmt.Println(array)
	fmt.Println(increment)

	// Mapas
	dictionary := map[string]int{
		"clave1": 1,
		"clave2": 2,
		"clave3": 3,
	}

	fmt.Println(dictionary)

	// estructura

	type Person struct {
		Nombre string
		Edad   int
	}

	// Crear una instancia de la estructura
	person := Person{Nombre: "Fainner", Edad: 29}
	fmt.Println(person)
}
