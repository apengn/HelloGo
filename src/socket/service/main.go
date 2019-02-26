package main

import (
	"net"
	"log"
	"fmt"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":8088")

	//ch := make(chan interface{})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("service start.....")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			break
		}
		go readData(conn)

	}

}

func readData(conn net.Conn) {

	b := make([]byte, 4096)

	for {
		cnt, err := conn.Read(b)

		if err != nil || cnt == 0 {
			conn.Close()
			return
		}
		readData := strings.TrimSpace(string(b[0:cnt]))

		inputs := strings.Split(readData, " ")
		switch inputs[0] {
		case "quit":

			conn.Close()
		case "ping":

			conn.Write([]byte("pong\n"))
		case "echo":

			echoStr := strings.Join(inputs[1:], " ")
			conn.Write([]byte(echoStr))
		default:
			fmt.Printf("Unsupported command: %s\n", inputs[0])
		}
	}
	fmt.Printf("Connection from %v closed. \n", conn.RemoteAddr())
}
