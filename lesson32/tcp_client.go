package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error connecting to server: ", err)
		return
	}

	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter message: ")
		message, _ := reader.ReadString('\n')

		conn.Write([]byte(message))

		response, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println("Error reading response: ", err)
			return
		}

		fmt.Println("Server response: ", response)
	}
}
