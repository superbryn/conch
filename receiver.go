package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type DiscoveryPackage struct {
	Username string `json:"username"`
	IP       string `json:"ip"`
	Payload  string `json:"payload"`
}

func main() {
	addr, _ := net.ResolveUDPAddr("udp4", ":9999")
	conn, _ := net.ListenUDP("udp4", addr)

	fmt.Println("Listing for peers")

	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, _ := conn.ReadFromUDP(buffer)

		var peer DiscoveryPackage
		err := json.Unmarshal(buffer[:n], &peer)

		if err != nil {
			break
		}

		fmt.Printf("--Peer Found--\n")
		fmt.Printf("Name: %s\n", peer.Username)
		fmt.Printf("IP: %s (Reported: %s)\n", remoteAddr.IP, peer.IP)
		fmt.Printf("Msg: %s\n\n", peer.Payload)
	}
}
