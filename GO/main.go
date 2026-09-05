package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var count = 0

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error", err)
			continue
		}
		count++
		fmt.Println("paisi ", count)
		_, _ = io.Copy(conn, conn)
		conn.Close()
	}

}
