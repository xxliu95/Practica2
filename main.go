package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/xxliu95/Practica2/Generic"
	"github.com/xxliu95/Practica2/MutexStruct"
)

func init() {
	fmt.Printf("Init main Practica 2\n")
}

func GoAddLists(a []int, b []int) chan int {
	c := make(chan int, 3)
	for i := range a {
		result := a[i] + b[i]
		c <- result
	}
	return c
}

func main() {
	// Ejercicio 1
	fmt.Printf("\nEjercicio 1\n")

	slice1 := []int{1, 2, 3, 4, 5, 6}
	f1 := func(n int) bool {
		return n%2 == 0
	}
	ej1 := Generic.AnyOf(slice1, f1)
	fmt.Printf("%v Any %%2 == 0 %t\n", slice1, ej1)

	// Equals
	slice2 := []int{1, 2, 3, 4, 5, 6}
	ej2 := Generic.Equal(slice1, slice2)
	fmt.Printf("%v = %v %t\n", slice1, slice2, ej2)

	// Ejercicio 2
	fmt.Printf("\nEjercicio 2\n")
	myStruct := MutexStruct.MyStruct{}
	myStruct.Increment()
	fmt.Printf("Increment by 1: %d\n", myStruct.Value())
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		if rand.Int()%2 == 0 {
			myStruct.Increment()
		} else {
			myStruct.Value()
		}
	}
	fmt.Printf("Increment x times: %d\n", myStruct.Value())
	myStruct.IncrementBy(6)
	fmt.Printf("Increment by 6: %d\n", myStruct.Value())

	// Ejericio 3
	fmt.Printf("\nEjercicio 3\n")
	list1 := []int{1, 2, 3}
	list2 := []int{4, 2, 6}
	ch := GoAddLists(list1, list2)
	close(ch)

	for val := range ch {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}
