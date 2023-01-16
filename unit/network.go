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
)

const (
	Byte = bits * 8
	KB   = 1024 * Byte
	MB   = 1024 * KB
	GB   = 1024 * MB
	TB   = 1024 * GB
	PB   = 1024 * TB
	EB   = 1024 * PB
)

func Speed2Str(speed float64) string {
	return Speed2bps(speed)
}

func Speed2bps(speed float64) string {
	if speed < 0 {
		return "——"
	} else if speed == 0 {
		return "0bps"
	} else if speed < Kb {
		return fmt.Sprintf("%.2fbps", speed)
	} else if speed < Mb {
		return fmt.Sprintf("%.2fKbps", speed/Kb)
	} else if speed < Gb {
		return fmt.Sprintf("%.2fMbps", speed/Mb)
	} else if speed < Tb {
		return fmt.Sprintf("%.2fGbps", speed/Gb)
	} else if speed < Pb {
		return fmt.Sprintf("%.2fTbps", speed/Tb)
	} else { // if speed < Eb
		return fmt.Sprintf("%.2fPbps", speed/Pb)
	}
}

func Speed2Bs(speed float64) string {
	if speed < 0 {
		return "——"
	} else if speed == 0 {
		return "0B/s"
	} else if speed < KB {
		return fmt.Sprintf("%.2fB/s", speed)
	} else if speed < MB {
		return fmt.Sprintf("%.2fKB/s", speed/KB)
	} else if speed < GB {
		return fmt.Sprintf("%.2fMB/s", speed/MB)
	} else if speed < TB {
		return fmt.Sprintf("%.2fGB/s", speed/GB)
	} else if speed < PB {
		return fmt.Sprintf("%.2fTB/s", speed/TB)
	} else { // if speed < EB
		return fmt.Sprintf("%.2fPB/s", speed/PB)
	}
}

func Flow2Str(fileSize int64) (size string) {
	return Size2Str(fileSize)
}

func Flow2B(fileSize int64) (size string) {
	return Size2B(fileSize)
}

func Size2Str(fileSize int64) string {
	return Size2B(fileSize)
}

func Size2b(fileSize int64) string {
	if fileSize < 0 {
		return "——"
	} else if fileSize < Kb {
		return fmt.Sprintf("%.2fb", float64(fileSize))
	} else if fileSize < Mb {
		return fmt.Sprintf("%.2fKb", float64(fileSize)/float64(Kb))
	} else if fileSize < Gb {
		return fmt.Sprintf("%.2fMb", float64(fileSize)/float64(Mb))
	} else if fileSize < Tb {
		return fmt.Sprintf("%.2fGb", float64(fileSize)/float64(Gb))
	} else if fileSize < Pb {
		return fmt.Sprintf("%.2fTb", float64(fileSize)/float64(Tb))
	} else { // if fileSize < Eb
		return fmt.Sprintf("%.2fPb", float64(fileSize)/float64(Pb))
	}
}

func Size2B(fileSize int64) string {
	if fileSize < 0 {
		return "——"
	} else if fileSize < KB {
		return fmt.Sprintf("%.2fB", float64(fileSize))
	} else if fileSize < MB {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(KB))
	} else if fileSize < GB {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(MB))
	} else if fileSize < TB {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(GB))
	} else if fileSize < PB {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(TB))
	} else { // if fileSize < EB
		return fmt.Sprintf("%.2fPB", float64(fileSize)/float64(PB))
	}
}
