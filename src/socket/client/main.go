package main

import (
	"net"
	"log"
	"bufio"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8088")

	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	buf := make([]byte, 1024)

	for {
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			return
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Fatal(err)
			return
		}
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Fatal(string(buf[0:cnt]))

	}
}
