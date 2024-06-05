package main

import (
    "fmt"
    "net"
    "bufio"
)

var clients []net.Conn

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error creating listener:", err)
        return
    }
    defer listener.Close()

    fmt.Println("TCP server started on :8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        clients = append(clients, conn)
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("Client connected:", conn.RemoteAddr().String())

    for {
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("Error reading message:", err)
            removeClient(conn)
            return
        }

        fmt.Println("Received:", message)
        message = message[:len(message)-1]
        broadcastMessage(conn, message)
    }
}

func broadcastMessage(sender net.Conn, message string) {
    for _, client := range clients {
        if client != sender {
            _, err := client.Write([]byte(message + " FROM " + sender.RemoteAddr().String() + "\n"))
            if err != nil {
                fmt.Println("Error broadcasting message:", err)
            }
        }
    }
}

func removeClient(conn net.Conn) {
    for i, client := range clients {
        if client == conn {
            clients = append(clients[:i], clients[i+1:]...)
            break
        }
    }
}
