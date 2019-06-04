package main

import (
	"bufio"
	"fmt"
	"net"
)

var mensaje string

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	send(msg)
}

func send(msg string) {
	con2, _ := net.Dial("tcp", "10.11.98.200:8000")
	defer con2.Close()
	fmt.Print("Mensaje: ")
	fmt.Fprint(con2, msg) //envia mensaje
}

func main() {
	//servidor
	ln, _ := net.Listen("tcp", "10.11.97.199:8000")
	defer ln.Close()
	con, _ := ln.Accept()
	handle(con)
}