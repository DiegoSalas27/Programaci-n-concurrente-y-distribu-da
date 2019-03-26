package main

import "fmt"

type Animal struct {
	name string
	weight float64
}

func WordCount(s string) map[string]int {
	return map[string]int{"x":1}
}

func main() {
	a := make(map[string]int)
	b := 20
	c := map[string]int{
		"veinte": 20,
		"treinta": 30,
	}
	a["test1"] = b
	d := map[string]Animal{
		"fido": Animal{name: "fido"}}

	fmt.Println(a, c, d)
	fmt.Println(WordCount("hella"))
}