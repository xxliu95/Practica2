package main

import (
	"fmt"

	"github.com/xxliu95/Practica2/Generic"
)

func init() {
	fmt.Printf("Init main Practica 2\n")
}

func main() {
	// Ejercicio 1
	// AnyOf
	slice1 := []int{1, 2, 3, 4, 5, 6}
	f1 := func(n int) bool {
		return n%2 == 0
	}
	ej1 := Generic.AnyOf(slice1, f1)
	fmt.Printf("%t\n", ej1)

	// Equals
	slice2 := []int{1, 2, 3, 4, 5, 6}
	ej2 := Generic.Equal(slice1, slice2)
	fmt.Printf("%t\n", ej2)
}
