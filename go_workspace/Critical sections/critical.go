package main

import (
	"fmt"
	"math/rand"
	"time"
)

// using dual channels

func f1(do chan bool, next chan bool) {
	<-do
	fmt.Println("f1: Executing critical section 1...")
	next <- true
	fmt.Println("f1: Executing something else...\n")
}

func f2(do chan bool, next chan bool) {
	<-do
	fmt.Println("f2: Executing critical section 2...")
	next <- true
	fmt.Println("f2: Executing something else...\n")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Dual channel program")
	cf1 := make(chan bool, 1)
	cf2 := make(chan bool, 1)
	cf1 <- true
	sum := 0
	for sum < 5 {
		go f1(cf1, cf2)
		time.Sleep(1000 * time.Millisecond)
		go f2(cf2, cf1)
		sum += sum
	}
}
