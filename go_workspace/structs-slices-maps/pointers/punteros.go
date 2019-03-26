package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func main() {
	a := 10
	ptr := &a

	fmt.Println(a, ptr, *ptr)

	var ptr2 *int

	ptr2 = new(int)

	*ptr2 = 20

	fmt.Println(ptr2, *ptr2)

	var a1 Animal
	a1.name = "Fido"
	a2 := Animal{"Rafa"}
	a3 := Animal{name: "Perico"}
	
	ptr3 := &a3
	ptr4 := &Animal{"Michi"}

	fmt.Println(a1, a2, a3, &ptr3, &ptr4, *ptr3, *ptr4)

	var animales []Animal 
	animales = append(animales, Animal{name: "Pablo"})

	fmt.Println(animales)
}