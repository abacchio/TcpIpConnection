package main

import (
	"fmt"
	"net"
)

func main() {
	server()
	//client()
}

func server() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("Start listen to localhost")
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		buffer := make([]byte, 1024)
		go func() {
			fmt.Printf("[Remote address]\n%s\n", conn.RemoteAddr())

			n, _ := conn.Read(buffer)
			fmt.Printf("[Message]\n%s", string(buffer[:n]))

			response := fmt.Sprintf("Hello,%s\n", conn.RemoteAddr())
			fmt.Println(response)
			conn.Write([]byte(response))

			conn.Close()
		}()
	}
}

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	message := fmt.Sprintf("Hello, %s\n", conn.RemoteAddr())
	conn.Write([]byte(message))

	response := make([]byte, 1024)
	n, _ := conn.Read(response)
	fmt.Println(string(response[:n]))
}
