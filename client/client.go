package main

import (
	"fmt"
	"log"
	"net"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func Connect() (username string, connection net.Conn) {
	fmt.Scanln(&username)
	connection, err := net.Dial(config.SERVER_TYPE, config.HOST+":"+config.PORT)
	if err != nil {
		panic(err)
	}
	_, err = connection.Write([]byte("Client connected: " + username))
	if err != nil {
		fmt.Println("Unable to send message to server.")
	}
	return
}

func Prompt() (message string) {
	fmt.Scanln(&message)
	return
}

func listen(connection net.Conn) {
	go func() {
		for {
			var buffer []byte
			connection.Read(buffer)
			fmt.Println(string(buffer))
		}
	}()
}

func main() {
	username, connection := Connect()
	listen(connection)
	for {
		message := Prompt()
		connection.Write([]byte(username + ": " + message))
	}
}
