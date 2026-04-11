package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	broadcastAddr = "255.255.255.255:9999"
	listenAddr    = ":9999"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// 1. Start the Receiver in the background
	// It listens for anyone shouting on port 9999
	go startReceiver(name)

	// 2. Setup the "Megaphone" (Broadcast Sender)
	addr, _ := net.ResolveUDPAddr("udp4", broadcastAddr)
	// Important: We use DialUDP but the OS knows 255.255.255.255 means "everyone"
	conn, _ := net.DialUDP("udp4", nil, addr)

	fmt.Printf("\n--- CONCH MESH ACTIVE [%s] ---\n", name)
	fmt.Println("Anything you type will be sent to EVERYONE in the lab.")
	fmt.Println("-------------------------------------------------------")

	for {
		fmt.Print("> ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		if msg != "" {
			fullMsg := fmt.Sprintf("%s: %s", name, msg)
			conn.Write([]byte(fullMsg))
		}
	}
}

func startReceiver(myName string) {
	addr, _ := net.ResolveUDPAddr("udp4", listenAddr)
	conn, _ := net.ListenUDP("udp4", addr)
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		message := string(buffer[:n])

		// Logic: If the message starts with our own name, ignore it 
		// so we don't see our own messages echoed back.
		if !strings.HasPrefix(message, myName+":") {
			fmt.Printf("\r%s\n> ", message)
		}
	}
}