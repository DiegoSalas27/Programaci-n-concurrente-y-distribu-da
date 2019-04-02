package main

import (
	"fmt"
	"strings"
)

func main() {
	primes := []int{1, 3, 5, 7, 11, 13}

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

	// Slice literals

	booleanos1 := [3]bool{true, true, false}
	booleanos2 := []bool{true, true, false}

	fmt.Println(booleanos1, booleanos2)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	//Slice defaults

	a := []int{2, 3, 5, 7, 11, 13}

	a = a[1:4]
	fmt.Println(a)

	a = a[:2]
	fmt.Println(a)

	a = a[1:]
	fmt.Println(a)

	//Slice length and capacity

	t := []int{2, 3, 5, 7, 11, 13}
	printSlice(t)

	//Slice the slice to give it a zero length.
	t = t[:0]
	printSlice(t)

	// Extend its length
	t = t[:4]
	printSlice(t)

	//Drop its first two values
	t = t[2:]
	printSlice(t)

	//Nil slices

	var b []int
	fmt.Println(b, len(b), cap(b))
	if b == nil {
		fmt.Println("Nil!")
	}

	//Creating a slice with make

	c := make([]int, 5)
	printSlice2("c", c)

	d := make([]int, 0, 5)
	printSlice2("d", d)

	e := d[:2]
	printSlice2("e", e)

	f := e[2:5]
	printSlice2("f", f)

	//Slices of slices

	//Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func printSlice(t []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(t), cap(t), t)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
