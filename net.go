package utils

import (
	"math/big"
	"net"
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
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return -1
	}

	i := big.NewInt(0)
	i.SetBytes(ip)
	return i.Int64()
	//
	// bits := strings.Split(ipStr, ".")
	// if len(bits) < 3 {
	// 	return -1
	// }
	//
	// b0, _ := strconv.ParseInt(bits[0], 10, 64)
	// b1, _ := strconv.ParseInt(bits[1], 10, 64)
	// b2, _ := strconv.ParseInt(bits[2], 10, 64)
	// b3, _ := strconv.ParseInt(bits[3], 10, 64)
	//
	// var sum int64
	//
	// sum += b0 << 24
	// sum += b1 << 16
	// sum += b2 << 8
	// sum += b3
	//
	// return sum
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
