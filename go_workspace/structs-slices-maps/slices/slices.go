package main

import "fmt"

func main() {
	primes := []int{1,3,5,7,11,13}

	var slice []int = primes[1:4]
	fmt.Println(len(slice), slice)

	//slice as an array reference

	names := [4]string{
		"Diego",
		"Manuel",
		"Miguel",
		"Enzo",
	}

	fmt.Println(names)

	slice1 := names[0:2]
	slice2 := names[1:3]

	fmt.Println(slice1, slice2)

	slice2[0] = "XXX"
	fmt.Println(slice1, slice2)
	fmt.Println(names)
}