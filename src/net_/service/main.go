package main

import (
	"net"
	"HelloGo/src/net_/model"
	"bufio"
	"sync"
	"io"
	"strings"
	"fmt"
)

type SocketListener struct {
	listen   net.Listener
	lock     *sync.Mutex
	handfunc map[string]func(rw *bufio.ReadWriter)
}

func main() {

	socketListener := &SocketListener{handfunc: make(map[string]func(rw *bufio.ReadWriter))}
	socketListener.lock=new(sync.Mutex)
	socketListener.handfunc["xxx"] = func(rw *bufio.ReadWriter) {
		fmt.Println(rw.ReadString('\n'))
		rw.WriteString("Thank YOU  \n")
		rw.Flush()
	}
	var err error
	socketListener.listen, err = net.Listen("tcp", model.Port)
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	for {
		fmt.Printf("accept\n")
		conn, err := socketListener.listen.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			continue
		}

		go socketListener.handleData(conn)
	}

}

func (s *SocketListener) handleData(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()
	for {

		k, err := rw.ReadString('\n')
		fmt.Println(k)
		switch {
		case err == io.EOF:
			fmt.Printf("Reached EOF -close this connect .\n",err)
			return
		case err != nil:
			fmt.Printf("\nError reading command .Get:,err:%v", err)
			return
		}
		k = strings.Trim(k, "\n")

		s.lock.Lock()
		fc, ok := s.handfunc[k]
		s.lock.Unlock()
		if ok {
			fc(rw)
		}
	}

}
