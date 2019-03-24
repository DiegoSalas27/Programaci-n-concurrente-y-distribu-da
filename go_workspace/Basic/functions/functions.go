package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func addShortened(x, y int) int {
	return x + y
}

func addString(x string, y string) string {
	return x + y
}

func main() {
	fmt.Println(add(10, 12))
	fmt.Println(addString("ho", "la"))
	fmt.Println(addShortened(10, 12))
}
