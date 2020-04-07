package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/Rule-BasedGO/drone"
	"github.com/Rule-BasedGO/utils"
)

func main() {
	statementChan := make(chan string)
	stopchan := make(chan struct{})

	setupDrone("big jet 345", statementChan, stopchan)

	fmt.Println(fmt.Sprintf("listening on port %s", os.Getenv("LISTEN_PORT")))
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("LISTEN_PORT")))
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleRequest(conn, statementChan)
	}
}

func setupDrone(name string, statementChan chan string, stopchan chan struct{}) {
	d := drone.CreateNewDrone(name)
	go d.ClassifyStatements(statementChan, stopchan)
}

func handleRequest(conn net.Conn, statementChan chan string) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	n := bytes.Index(buf, []byte{0})
	message := string(buf[:n])
	fmt.Println("message received: ", message)
	cleaned := utils.WordToNum(strings.ToLower(message))
	statementChan <- cleaned
	// Close the connection when you're done with it.
	conn.Close()
}
