package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
	fmt.Print("Enter username: ")
	// fmt.Scanln(&username)
	username = "Owen"
	connection, err := net.Dial(config.SERVER_TYPE, config.HOST+":"+config.PORT)
	if err != nil {
		panic(err)
	}
	_, err = connection.Write([]byte(username + " connected"))
	if err != nil {
		fmt.Println("Unable to send message to server.")
	}
	return
}

func Prompt() (message string) {
	fmt.Print("Type message: ")
	fmt.Scanln(&message)
	return
}

func Listen(connection net.Conn) {
	for {
		var buffer []byte
		connection.Read(buffer)
		if len(buffer) > 0 {
			fmt.Println(string(buffer))
		}
	}
}

func main() {
	log.Println("Configuring client...")
	Configure(os.Args)
	log.Println("Client configured")
	log.Println("Connecting to server...")
	username, connection := Connect()
	go Listen(connection)
	for {
		message := Prompt()
		connection.Write([]byte(username + ": " + message))
	}
}
