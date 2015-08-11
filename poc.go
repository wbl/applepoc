package main

import "fmt"
import "log"
import "net"
import "smallgroup"

func main() {
	cert, err := smallgroup.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Fatal("error ", err)
	}
	config := smallgroup.Config{Certificates: []smallgroup.Certificate{cert},
		CipherSuites: []uint16{0xc02f, 0xc014},
		
	}
	listener, err := smallgroup.Listen("tcp", "127.0.0.1:1443", &config)
	if err != nil {
		log.Fatal("error ", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error on connection!")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	conn.Read(make([]byte, 40, 40))
	fmt.Printf("Connection")
	fmt.Fprintf(conn, "HTTP/1.0 500 Internal Server Error\r\n")
	fmt.Fprintf(conn, "Content-type: text/plain;\r\n\r\n")
	fmt.Fprintf(conn, "Internal Error!")
	conn.Close()
}
