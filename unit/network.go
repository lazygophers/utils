package unit

import (
	"fmt"
	"strconv"
)

const (
	Byte = 1
	KB   = 1024 * Byte
	MB   = 1024 * KB
	GB   = 1024 * MB
	TB   = 1024 * GB
	PB   = 1024 * TB
	EB   = 1024 * PB
)

const (
	Bit = 8 * Byte
	Kb  = 1024 * Bit
	Mb  = 1024 * Kb
	Gb  = 1024 * Mb
	Tb  = 1024 * Gb
	Pb  = 1024 * Tb
	Eb  = 1024 * Pb
)

func FormatSpeedWithUnit(speed float64) (string, string) {
	if speed <= 0 {
		return "——", "——"
	} else if speed < Kb/8 {
		return strconv.FormatFloat(speed*8, 'f', 2, 64), "bps"
	} else if speed < Mb/8 {
		return strconv.FormatFloat(speed*8/Kb, 'f', 2, 64), "Kbps"
	} else if speed < Gb/8 {
		return strconv.FormatFloat(speed*8/Mb, 'f', 2, 64), "Mbps"
	} else if speed < Tb/8 {
		return fmt.Sprintf("%.2f", speed*8/Gb), "Gbps"
	} else if speed < Pb/8 {
		return strconv.FormatFloat(speed*8/Tb, 'f', 2, 64), "Tbps"
	} else { // if speed < Eb
		return strconv.FormatFloat(speed*8/Pb, 'f', 2, 64), "Pbps"
	}
}

func FormatSpeed(speed float64) string {
	return Format2bps(speed)
}

func Format2bps(speed float64) string {
	if speed <= 0 {
		return "——"
	} else if speed < Kb/8 {
		return fmt.Sprintf("%.2fbps", speed*8)
	} else if speed < Mb/8 {
		return fmt.Sprintf("%.2fKbps", speed*8/Kb)
	} else if speed < Gb/8 {
		return fmt.Sprintf("%.2fMbps", speed*8/Mb)
	} else if speed < Tb/8 {
		return fmt.Sprintf("%.2fGbps", speed*8/Gb)
	} else if speed < Pb/8 {
		return fmt.Sprintf("%.2fTbps", speed*8/Tb)
	} else { // if speed < Eb
		return fmt.Sprintf("%.2fPbps", speed*8/Pb)
	}
}

func Format2Bs(speed float64) string {
	if speed <= 0 {
		return "——"
	} else if speed < Kb {
		return fmt.Sprintf("%.2fB/s", speed)
	} else if speed < Mb {
		return fmt.Sprintf("%.2fKB/s", speed/Kb)
	} else if speed < Gb {
		return fmt.Sprintf("%.2fMB/s", speed/Mb)
	} else if speed < Tb {
		return fmt.Sprintf("%.2fGB/s", speed/Gb)
	} else if speed < Pb {
		return fmt.Sprintf("%.2fTB/s", speed/Tb)
	} else { // if speed < EB
		return fmt.Sprintf("%.2fPB/s", speed/Pb)
	}
}

func FormatSize(fileSize int64) string {
	return Format2B(fileSize)
}

func Format2b(fileSize int64) string {
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

func Format2B(fileSize int64) string {
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
