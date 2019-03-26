package main

import (
	"fmt"
	"time"
	"math/rand"
)

var n int = 0

func pausa() {
	v := rand.Intn(20) 
	time.Sleep(time.Duration(v)*time.Millisecond)
}

func p() {
	var temp int
	for i:=0; i < 10; i++ {
		pausa()
		temp = n
		pausa()
		n = temp + 1
	}
}

func q() {
	var temp int
	for i:=0; i < 10; i++ {
		pausa()
		temp = n
		pausa()
		n = temp + 1
	}
}


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	go p()
	go q()
	time.Sleep(300*time.Millisecond)
	fmt.Println(n)
}