package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2) //los canales son un mecanismo de comunicacion entre procesos go y suelen proteger secciones criticas
	go producer(1, ch) //no se puede leer y escribir un canal en un solo proceso. a menos que utilicemos canal con buffer
	go consumer(2, ch)

	time.Sleep(time.Second)
}

func producer(id int, ch chan string) {
	c := 0
	for {
		ch <- fmt.Sprintf("producto %d producido por %d", c, id) 
		c++
	}
}

func consumer(id int, ch chan string) { //si el canal esta vacio, bloquea este proceso. Debe esperar a que sea escrito
	for {
		fmt.Printf("Consumidor %d usando %s\n", id, <-ch)
	}
}