package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msg(str string) {
	dur := time.Duration(rand.Intn(100) + 50)
	time.Sleep(dur * time.Nanosecond)
	fmt.Println(str)
}

var turn int = 1
var end chan bool

func P() {
	for i := 0; i < 10; i++ {
		msg("P: Mensaje 1 Sección no crítica")
		msg("P: Mensaje 2 Sección no crítica")
		msg("P: Mensaje 3 Sección no crítica")
		for turn != 1 {

		}
		msg("P: Mensaje 1 Sección  CRÍTICA")
		msg("P: Mensaje 2 Sección  CRÍTICA")
		msg("P: Mensaje 3 Sección  CRÍTICA")
		turn = 2
	}
	end <- true
}

func Q() {
	for i := 0; i < 10; i++ {
		msg("Q: Mensaje 1 Sección no crítica")
		msg("Q: Mensaje 2 Sección no crítica")
		msg("Q: Mensaje 3 Sección no crítica")
		for turn != 2 {

		}
		msg("Q: Mensaje 1 Sección  CRÍTICA")
		msg("Q: Mensaje 2 Sección  CRÍTICA")
		msg("Q: Mensaje 3 Sección  CRÍTICA")
		turn = 1
	}
	end <- true
}

func main() {
	end = make(chan bool)
	go P()
	go Q()
	<-end
	<-end
}
