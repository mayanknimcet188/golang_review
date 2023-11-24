package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// take port number as flag argument, default to 8000
	portNumber := flag.Int("port", 8000, "port number of the server, defaults to 8000")
	flag.Parse()
	server := fmt.Sprint("localhost:", *portNumber)
	listener, err := net.Listen("tcp", server)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// each connection can run on its own individual go routine, hence mutliple clients can listen
		// to the clock at the same time
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
