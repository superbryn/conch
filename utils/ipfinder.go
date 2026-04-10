package utils

import (
	"fmt"
	"net"
)

func IpAddrsFinder() string {
    interfaces, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }

    for _, addr := range interfaces {

        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }

    return ""
}