package unit

import (
	"fmt"
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

func FormatSpeed(speed float64) string {
	return Format2bps(speed)
}

func Format2bps(speed float64) string {
	if speed <= 0 {
		return "——"
	} else if speed < Kb/8 {
		return fmt.Sprintf("%.2f bps", speed*8)
	} else if speed < Mb/8 {
		return fmt.Sprintf("%.2f Kbps", speed*8/Kb)
	} else if speed < Gb/8 {
		return fmt.Sprintf("%.2f Mbps", speed*8/Mb)
	} else if speed < Tb/8 {
		return fmt.Sprintf("%.2f Gbps", speed*8/Gb)
	} else if speed < Pb/8 {
		return fmt.Sprintf("%.2f Tbps", speed*8/Tb)
	} else if speed < Eb/8 { // 确保在 Eb 范围内正确处理
		return fmt.Sprintf("%.2f Pbps", speed*8/Pb)
	} else { // 超过 Eb 的值直接返回最大单位
		return fmt.Sprintf("%.2f Ebps", speed*8/Eb)
	}
}

func Format2Bs(speed float64) string {
	if speed <= 0 {
		return "——"
	} else if speed < Kb {
		return fmt.Sprintf("%.2f B/s", speed)
	} else if speed < Mb {
		return fmt.Sprintf("%.2f KB/s", speed/Kb)
	} else if speed < Gb {
		return fmt.Sprintf("%.2f MB/s", speed/Mb)
	} else if speed < Tb {
		return fmt.Sprintf("%.2f GB/s", speed/Gb)
	} else if speed < Pb {
		return fmt.Sprintf("%.2f TB/s", speed/Tb)
	} else { // if speed < EB
		return fmt.Sprintf("%.2f PB/s", speed/Pb)
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
		return fmt.Sprintf("%.2f KB", float64(fileSize)) // 修改: 将 "%.2fB" 改为 "%.2f KB"
	} else if fileSize < MB {
		return fmt.Sprintf("%.2f KB", float64(fileSize)/float64(KB))
	} else if fileSize < GB {
		return fmt.Sprintf("%.2f MB", float64(fileSize)/float64(MB))
	} else if fileSize < TB {
		return fmt.Sprintf("%.2f GB", float64(fileSize)/float64(GB))
	} else if fileSize < PB {
		return fmt.Sprintf("%.2f TB", float64(fileSize)/float64(TB))
	} else if fileSize < EB { // 添加对 EB 单位的判断
		return fmt.Sprintf("%.2f PB", float64(fileSize)/float64(PB))
	} else { // 确保在 EB 范围内正确处理
		return fmt.Sprintf("%.2f EB", float64(fileSize)/float64(EB))
	}
}
