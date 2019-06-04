package main

import (
	"bufio"
	"fmt"
	"net"
)

var n int

func handle(con net.Conn, id int) {
	defer con.Close()
	r := bufio.NewReader(con)
	for {
		msg, _ := r.ReadString('\n')
		fmt.Fprint(con, "(%d) %s", n, msg)
		fmt.Printf("Con%d: (%d) %s", id, n, msg)
		temp := n
		n = temp + 1
		if len(msg) == 0 || msg[0] == 'x' {
			break
		}
	}
	fmt.Printf("Con%d cerrada!\n", id)
}

func main() {
	ln, _ := net.Listen("tcp", "10.11.97.199:8000")
	defer ln.Close()
	cont := 0
	for {
		con, _ := ln.Accept()
		go handle(con, cont)
		cont++
	}
}