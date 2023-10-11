package utils

import (
	"encoding/binary"
	"net"
)

func IPv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}

	return ""
}

func IPv4Int() uint32 {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return 0
	}

	for _, addr := range addrs {
		if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				return binary.BigEndian.Uint32(ip.IP.To4())
			}
		}
	}

	return 0
}
