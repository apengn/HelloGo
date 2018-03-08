package main

import (
	"bufio"
	"log"
	"net"
	"sync"
	"io"
	"strings"
	"encoding/gob"
	"errors"
	"fmt"
	"strconv"
	"flag"
)

type complexData struct {
	N int
	S string
	M map[string]int
	P []byte
	C *complexData
}

const Port = ":61000"

func Open(addr string) (*bufio.ReadWriter, error) {

	log.Println("Dial  ", addr)
	//Dial函数和服务端建立连接：
	conn, erro := net.Dial("tcp", addr)

	if erro != nil {
		return nil, erro
	}
	//把连接转换为输出输入流
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil

}

type HandleFunc func(writer *bufio.ReadWriter)

type Endpoint struct {
	listener net.Listener
	handle   map[string]HandleFunc
	m        sync.RWMutex
}

func NewEndpoint() *Endpoint {
	return &Endpoint{handle: map[string]HandleFunc{}}
}

//添加处理函数的方法
func (e *Endpoint) AddHandleFunc(name string, f HandleFunc) {
	e.m.Lock()
	e.handle[name] = f
	e.m.Unlock()
}

//监听是否有请求过来
func (e *Endpoint) Listen() (err error) {

	//监听tcp
	e.listener, err = net.Listen("tcp", Port)

	if err != nil {
		return
	}
	log.Println("Listen on", e.listener.Addr().String())

	for {
		log.Println("Accept a connection request.")
		//等待客户端发送请求  （阻赛）
		conn, err := e.listener.Accept()
		if err != nil {
			log.Println("Failed accepting a connection request:")
			continue
		}
		log.Println("Handle incoming messages.")
		go e.HandleMessages(conn)
	}
}

func (e *Endpoint) HandleMessages(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	defer conn.Close()

	for {
		log.Print("Receive command ")

		cmd, err := rw.ReadString('\n')

		switch {
		case err == io.EOF:
			log.Println("Reached EOF -close this connection.\n ----")
			return
		case err != nil:
			log.Println("\nError reading command. Got:"+cmd+"\n", err)
			return
		}
		cmd = strings.Trim(cmd, "\n")
		log.Println(cmd, " ' ")
		e.m.RLock()

		handleCommand, ok := e.handle[cmd]

		e.m.RUnlock()
		if !ok {
			log.Println("Command '" + cmd + " is not registered.")
			return
		}
		handleCommand(rw)
	}

}

func HandleStrings(rw *bufio.ReadWriter) {
	log.Println("Receive String message:")
	s, err := rw.ReadString('\n')
	if err != nil {
		log.Println("Cannot read from connection .\n", err)
	}

	s = strings.Trim(s, "\n  ")
	log.Println(s)
	_, err = rw.WriteString("Thank you .\n")
	if err != nil {
		log.Println("Cannot write to connection.\n", err)
	}
	//要Flush()在写入后调用，以便所有数据都被转发到底层网络连接
	err = rw.Flush()
	if err != nil {
		log.Println("Flush failed.", err)
	}
}

//利用gob反序列化本地的struct对象
func HandleGob(rw *bufio.ReadWriter) {

	log.Println("Receive Gob data:")

	var data complexData
	//通过gob  把rw 解析成一个Struct
	dec := gob.NewDecoder(rw)

	err := dec.Decode(&data)

	if err != nil {
		log.Println("Error decoding GOB data:", err)
		return
	}

	log.Printf("Outer complexData struct: \n%#v\n", data)
	log.Printf("Inner complexData struct: \n%#v\n", data.C)

}

func client(ip string) error {

	testStruct := complexData{
		N: 23,
		S: "string data",
		M: map[string]int{"one": 1, "two": 2, "three": 3},
		P: []byte("abc"),
		C: &complexData{N: 256, S: "Recursive structs? Piece of  cake!", M: map[string]int{"01": 1, "10": 2, "11": 3}},
	}
	//获取连接之后    像输入/输出流一样处理新的连接
	rw, err := Open(ip + Port)

	if err != nil {
		return errors.New(fmt.Sprintln(err, "Client:Failed to open connection to ", ip, Port))
	}
	log.Println("Send the string request.")

	n, err := rw.WriteString("STRING\n")

	if err != nil {
		return errors.New(fmt.Sprintln(err, "Could not send the STRING request (", strconv.Itoa(n), "bytes written)"))
	}
	n, err = rw.WriteString("Additional data.\n")

	if err != nil {
		return errors.New(fmt.Sprintln(err, "Could not send additional STRING data(", strconv.Itoa(n), "bytes written)"))
	}

	log.Println("Flush the buffer.")
	//要Flush()在写入后调用，以便所有数据都被转发到底层网络连接
	err = rw.Flush()

	if err != nil {
		return errors.New(fmt.Sprintln(err, "Flush failed"))
	}
	log.Println("Read the reply.")
	//等待服务回复
	response, err := rw.ReadString('\n')

	if err != nil {
		return errors.New(fmt.Sprintln(err, "Client:Failed to read the reply:", response))
	}

	log.Println("STRING request:got a response:", response)

	log.Println("Send a struct as GOB:")

	log.Printf("Outer complexData struct: \n%#v\n", testStruct)
	log.Printf("Outer complexData struct: \n%#v\n", testStruct.C)

	//gob包实现的序列化struct对象保存到本地
	//	务必注意的是golang序列化有个小坑,就是struct里的字段必须要可导出也就是首字母大写
	//通过输入输出流  返回一个新的编码器
	enc := gob.NewEncoder(rw)

	n, err = rw.WriteString("GOB\n")

	if err != nil {
		return errors.New(fmt.Sprintln(err, "Could not write GOB data ("+strconv.Itoa(n)+" bytes written)"))
	}
	//设置要编码的Struct  (也就是序列化对象)
	err = enc.Encode(testStruct)
	if err != nil {
		return errors.New(fmt.Sprintln(err, "Encode failed for struct: %#v", testStruct))
	}
	//要Flush()在写入后调用，以便所有数据都被转发到底层网络连接
	err = rw.Flush()
	if err != nil {
		return errors.New(fmt.Sprintln(err, "Flush failed."))
	}
	return nil
}
func server() error {
	endpoint := NewEndpoint()
	//添加处理string的方法
	endpoint.AddHandleFunc("STRING", HandleStrings)
	//添加通过gob处理Struct的方法
	endpoint.AddHandleFunc("GOB", HandleGob)

	return endpoint.Listen()
}

func main() {
	//命令行参数解析  :命令：go run dome1.go -connect localhost
	connect := flag.String("connect", "", "IP address of process to join. If empty, go into listen mode.")
	flag.Parse()
	if *connect != "" {
		err := client(*connect)
		if err != nil {
			log.Println("Error:", err)
		}
		log.Println("Client done.")
		return
	}
	err := server()
	if err != nil {
		log.Println("Error:", err)
	}

	log.Println("Server done.")
}
func init() {
	log.SetFlags(log.Lshortfile)
}
