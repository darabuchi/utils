package utils

import (
	"net"
	"time"
)

func TCPPing(address string) (time.Duration, error) {
	return DialTimeout("tcp", address, time.Second*3)
}

func UDPPing(address string) (time.Duration, error) {
	return DialTimeout("udp", address, time.Second*3)
}

func DialTimeout(network, addr string, timeout time.Duration) (time.Duration, error) {
	start := time.Now()
	conn, err := net.DialTimeout(network, addr, timeout)
	if err != nil {
		return -1, err
	}
	_ = conn.Close()

	return time.Since(start), nil
}
