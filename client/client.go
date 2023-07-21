package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var reader bufio.Reader = *bufio.NewReader(os.Stdin)

func Connect() (username string, connection net.Conn) {
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
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
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return
}

func Listen(connection net.Conn) {
	for {
		var buffer []byte = make([]byte, 100)
		connection.Read(buffer)
		fmt.Println(string(buffer))
	}
}

func main() {
	log.Println("Configuring client...")
	Configure(os.Args)
	log.Println("Client configured")
	log.Println("Connecting to server...")
	username, connection := Connect()
	log.Println("Connected to server. Chat is now open.")
	go Listen(connection)
	for {
		message := Prompt()
		connection.Write([]byte(username + ": " + message))
	}
}
