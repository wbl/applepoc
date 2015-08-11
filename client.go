package main

import "fmt"
import "crypto/tls"

func main() {
	var config =tls.Config{InsecureSkipVerify:true}
	conn, err :=tls.Dial("tcp", "127.0.0.1:1443", &config)
	if err != nil {
		fmt.Printf("error")
	} else {
		fmt.Printf("succes!")
	}
	fmt.Fprintf(conn, "hi!")
}
