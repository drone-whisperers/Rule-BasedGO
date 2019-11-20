package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Rule-BasedGO/drone"
)

func main() {
	statementChan := make(chan string)
	stopchan := make(chan struct{})

	go setupDrone("other jet 345", statementChan, stopchan)
	setupDrone("big jet 345", statementChan, stopchan)

	fmt.Print("Input stream now open: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			close(stopchan)
			return
		}
		statementChan <- strings.ToLower(text)
	}
}

func setupDrone(name string, statementChan chan string, stopchan chan struct{}) {
	d := drone.CreateNewDrone(name)
	go d.ClassifyStatements(statementChan, stopchan)
}
