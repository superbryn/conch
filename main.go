package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type DiscoveryPacket struct {
	Username string `json:"username"`
	IP       string `json:"ip"`
	Payload  string `json:"payload"`
}

const (
	port          = ":9999"
	broadcastAddr = "255.255.255.255:9999"
	myUsername    = "Neeraj"
)

func main() {
	fmt.Printf("--- Starting Conch P2P Discovery [%s] ---\n", myUsername)

	go startReceiver()

	startSender()
}

func startReceiver() {
	addr, _ := net.ResolveUDPAddr("udp4", port)
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("Receiver Error:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 2048)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		var peer DiscoveryPacket
		if err := json.Unmarshal(buffer[:n], &peer); err != nil {
			continue 
		}

		if peer.Username != myUsername {
			fmt.Printf("\n[Peer Online] %s (%s) - %s", peer.Username, remoteAddr.IP, peer.Payload)
			fmt.Print("\n> ") 
		}
	}
}

func startSender() {
	dest, _ := net.ResolveUDPAddr("udp4", broadcastAddr)
	conn, err := net.DialUDP("udp4", nil, dest)
	if err != nil {
		fmt.Println("Sender Error:", err)
		return
	}
	defer conn.Close()

	packet := DiscoveryPacket{
		Username: myUsername,
		IP:       "Local", 
		Payload:  "Looking for peers!",
	}

	for {
		jsonData, _ := json.Marshal(packet)
		_, err := conn.Write(jsonData)
		if err != nil {
			fmt.Println("Broadcast error:", err)
		}
		
		// Shout every 5 seconds
		time.Sleep(5 * time.Second)
	}
}