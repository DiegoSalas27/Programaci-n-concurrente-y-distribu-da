package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(" No se pudo iniciar servidor")
		return
	}
	fmt.Println("Servidor listo y corriendo")
	con, err := ln.Accept()
	if err != nil {
		fmt.Println(" No se pudo conectar")
		return
	}
	fmt.Println("Conexion establecida")
	r := bufio.NewReader(con)
	msg, err :=  r.ReadString('\n')
	if err != nil {
		fmt.Println(" No se pudo leer mensaje")
		return
	} else {
		fmt.Println("Recibido: ", msg)
	}
}