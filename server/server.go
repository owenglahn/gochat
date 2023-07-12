package main

import (
	"log"
	"net"
)

var CONNECTIONS []net.Conn

func SendToAllClients(msg []byte) {
	for _, conn := range CONNECTIONS {
		conn.Write(msg)
	}
}

func ListenToClient(conn net.Conn) {
	for {
		var buffer []byte
		conn.Read(buffer)
		if string(buffer) == "DISCONNECT" {
			return
		}
		if len(buffer) > 0 {
			SendToAllClients(buffer)
		}
	}
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
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		CONNECTIONS = append(CONNECTIONS, conn)
		go ListenToClient(conn)
	}
}
