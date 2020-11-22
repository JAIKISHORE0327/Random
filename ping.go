package main

import (
	"fmt"
)

func main() {
	pingChan := make(chan string)
	pongChan := make(chan string)

	go printer(pongChan)
	go pinger(pingChan)
	go ponger(pingChan, pongChan)

	// Need to block until all complete
	var input string
	fmt.Scanln(&input)
}

func pinger(pingChan chan<- string) {
	// Send a ping
	fmt.Println("pinger sending \"ping\"")
	pingChan <- "ping"
}

func ponger(pingChan <-chan string, pongChan chan<- string) {
	// receive the ping
	fmt.Println("ponger received", <-pingChan)

	// respond with a pong
	fmt.Println("ponger replying with \"pong\"")
	pongChan <- "pong"
}

func printer(pongChan <-chan string) {
	// read the pong
	fmt.Println("printer received", <-pongChan)
}