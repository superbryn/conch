package main

import (
	"encoding/json"
	"net"
	"time"
	"github.com/superbryn/conch/utils"
)

type DiscoveryPackage struct{
	Username string `json:"username"`
	IP string `json:"ip"`
	Payload string `json:"payload"`
}

func main(){
	dest, _ := net.ResolveUDPAddr("udp4","255.255.255.255:6767")
	conn, _ := net.DialUDP("udp4",nil,dest)
	ip_address := utils.IpAddrsFinder()
	
	packet := DiscoveryPackage{
		Username: "Neeraj",
		IP: ip_address,
		Payload: "Conch Discovery",
	}
	
	for {
		jsonData, _ := json.Marshal(packet)
		
		conn.Write(jsonData)
		
		time.Sleep(5 * time.Second)
	}
}