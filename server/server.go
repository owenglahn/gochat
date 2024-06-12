package main

import (
	"log"
	"net"
	"strings"
)

var CONNECTIONS []net.Conn

func SendToAllClients(msg []byte) {
	for _, conn := range CONNECTIONS {
		conn.Write(msg)
	}
}

func ListenToClient(conn net.Conn) {
	for {
		var buffer []byte = make([]byte, 100)
		conn.Read(buffer)
		if strings.Contains(string(buffer), "DISCONNECT") {
			username := strings.Split(string(buffer), ":")[0]
			msg := username + " disconnected"
			log.Println(msg)
			SendToAllClients([]byte(msg))
			conn.Close()
			return
		}
		log.Println("NEW MESSAGE: " + string(buffer))
		SendToAllClients(buffer)
	}
}

func shutdown(listener net.Listener) {
	log.Println("Shutting down server...")
	listener.close()
}	

func main() {
	log.Println("Starting server...")
	log.Println("Configuring server...")
	Configure()
	log.Println("Server configured.")
	listener, err := net.Listen(config.SERVER_TYPE, config.HOST+":"+config.PORT)
	log.Println("Listening on " + config.HOST + ":" + config.PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		log.Println("Client connected. ID: " + conn.LocalAddr().String())
		CONNECTIONS = append(CONNECTIONS, conn)
		go ListenToClient(conn)
	}
}
