package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println(add("he", "la"))
	fmt.Scanln()
}

func add(x string, y string) string {
	return x + y
}
