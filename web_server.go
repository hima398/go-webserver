package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
)

type WebServer struct {
	port    uint16
	gateway string
}

func NewWebServer(port uint16, gateway string) *WebServer {
	return &WebServer{port, gateway}
}

func (ws *WebServer) Run() {
	fmt.Println("Run Web Server.")
	l, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println("already used 80 port.")
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			// handle error
		}
		go func(conn net.Conn) {
			content, err := Read(conn)
			if err != nil {
				// handle error
			}
			fmt.Printf("%s\n", content)
		}(conn)
	}
}

func Read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	var buffer bytes.Buffer
	for {
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		buffer.Write(ba)
		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}
