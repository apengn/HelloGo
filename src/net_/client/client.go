package main

import (
	"net"
	"log"
	"flag"
	"HelloGo/src/net_/model"
	"bufio"
	"fmt"
)

func main() {

	ipaddress := flag.String("connect", "localhost"+model.Port, "")
	flag.Parse()

	conn, err := net.Dial("tcp", *ipaddress)

	if err != nil {
		log.Fatal("connect faile",err)
		return
	}

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	rw.WriteString("xxx\n")
	rw.WriteString("xxxxxxxxxxxxxxxxxxxxxxxxxxxx\n")

	rw.Flush()


	fmt.Println(rw.ReadString('\n'))
	conn.Close()

}
