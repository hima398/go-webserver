package main

import (
	"fmt"
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
	l, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		fmt.Println("already used 80 port.")
	}
	defer l.Close()
}
