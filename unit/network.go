package unit

import (
	"fmt"
)

const (
	bits = 1
	Kb   = 1024 * bits
	Mb   = 1024 * Kb
	Gb   = 1024 * Mb
	Tb   = 1024 * Gb
	Pb   = 1024 * Tb
	Eb   = 1024 * Pb
	Zb   = 1024 * Eb
	Yb   = 1024 * Zb
)

const (
	Byte = bits * 8
	KB   = 1024 * Byte
	MB   = 1024 * KB
	GB   = 1024 * MB
	TB   = 1024 * GB
	PB   = 1024 * TB
	EB   = 1024 * PB
	ZB   = 1024 * EB
	YB   = 1024 * ZB
)

func Speed2Str(speed float64) string {
	var speedStr string
	if speed == 0 {
		speedStr = "0bps"
	} else if speed < 1024 {
		speedStr = fmt.Sprintf("%.2fbps", speed*8)
	} else if speed < (1024 * 128) {
		speedStr = fmt.Sprintf("%.2fKbps", speed/128)
	} else if speed < (1024 * 1024 * 128) {
		speedStr = fmt.Sprintf("%.2fMbps", speed/(128*1024))
	} else if speed < (1024 * 1024 * 1024 * 128) {
		speedStr = fmt.Sprintf("%.2fGbps", speed/(128*1024*1024))
	} else {
		speedStr = fmt.Sprintf("%.2fTbps", speed/(128*1024*1024*1024))
	}

	return speedStr
}

func Flow2Str(fileSize uint64) (size string) {
	if fileSize < 1024 {
		// return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { // if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}
