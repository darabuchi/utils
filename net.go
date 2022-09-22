package utils

import (
	"net"
	"strconv"
	"strings"
)

func IsIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsIpV4(ip string) bool {
	IP := net.ParseIP(ip)
	if IP == nil {
		return false
	}

	return IP.To4() != nil
}

func IsIpV6(ip string) bool {
	IP := net.ParseIP(ip)
	if IP == nil {
		return false
	}

	return IP.To16() != nil
}

func Ip2Int(ipStr string) int64 {
	bits := strings.Split(ipStr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

func Int2Ip(ip int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ip & 0xFF)
	bytes[1] = byte((ip >> 8) & 0xFF)
	bytes[2] = byte((ip >> 16) & 0xFF)
	bytes[3] = byte((ip >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func IsLocalIp(ip string) bool {
	i := net.ParseIP(ip)
	if i == nil {
		return false
	}

	if i.To4() == nil {
		return false
	}

	inputIpNum := Ip2Int(ip)

	innerIpF := Ip2Int("127.255.255.255")
	if inputIpNum>>24 == innerIpF>>24 {
		return true
	}

	innerIpC := Ip2Int("192.168.255.255")
	if inputIpNum>>16 == innerIpC>>16 {
		return true
	}

	innerIpB := Ip2Int("172.16.255.255")
	if inputIpNum>>20 == innerIpB>>20 {
		return true
	}

	innerIpA := Ip2Int("10.255.255.255")
	if inputIpNum>>24 == innerIpA>>24 {
		return true
	}

	innerIpD := Ip2Int("100.64.255.255")
	if inputIpNum>>22 == innerIpD>>22 {
		return true
	}

	return false
}
