package main

import "fmt"

func swap(x, y, z string) (string, string, string) {
	return x, y, z
}

func main() {
	a, b, c := swap("hello", "world", "holitas")
	fmt.Println(a, b, c)
}
