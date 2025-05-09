package unit

import (
	"fmt"
	"sync"
)

const (
	Byte = 1
	KB   = 1024 * Byte
	MB   = 1024 * KB
	GB   = 1024 * MB
	TB   = 1024 * GB
	PB   = 1024 * TB
	EB   = 1024 * PB

	Bit = 8 * Byte
	Kb  = 1024 * Bit
	Mb  = 1024 * Kb
	Gb  = 1024 * Mb
	Tb  = 1024 * Gb
	Pb  = 1024 * Tb
	Eb  = 1024 * Pb
)

// 全局 units 切片，减少内存分配
var (
	bpsUnits = []struct {
		value float64
		unit  string
	}{
		{Kb, "Kbps"},
		{Mb, "Mbps"},
		{Tb, "Tbps"},
		{Pb, "Pbps"},
		{Eb, "Ebps"},
		{float64(^uint64(0)), "Ebps"},
	}

	sizeUnits = []struct {
		value float64
		unit  string
	}{
		{float64(KB), "KB"},
		{float64(MB), "MB"},
		{float64(GB), "GB"},
		{float64(TB), "TB"},
		{float64(PB), "PB"},
		{float64(EB), "EB"},
		{float64(^uint64(0)), "EB"},
	}

	rateUnits = []struct {
		value float64
		unit  string
	}{
		{float64(KB), "KB/s"},
		{float64(MB), "MB/s"},
		{float64(GB), "GB/s"},
		{float64(TB), "TB/s"},
		{float64(PB), "PB/s"},
		{float64(EB), "EB/s"},
		{float64(^uint64(0)), "EB/s"},
	}
)

// sync.Pool 缓存 fmt.Sprintf 的结果
var formatPool = sync.Pool{
	New: func() interface{} {
		return &struct {
			result string
		}{}
	},
}

// Format2bps 格式化比特率，支持从 Kbps 到 Ebps 的单位转换
func Format2bps(speed float64) string {
	if speed <= 0 {
		return "——"
	}
	for _, u := range bpsUnits {
		if speed < u.value {
			return fmt.Sprintf("%.2f %s", speed/u.value, u.unit)
		}
	}
	return fmt.Sprintf("%.2f Ebps", speed/Eb)
}

// Format2B 格式化字节数，支持从 KB 到 EB 的单位转换
func Format2B(fileSize int64) string {
	if fileSize < 0 {
		return "——"
	}
	for _, u := range sizeUnits {
		if float64(fileSize) < u.value {
			return fmt.Sprintf("%.2f %s", float64(fileSize)/u.value, u.unit)
		}
	}
	return fmt.Sprintf("%.2f EB", float64(fileSize)/float64(EB))
}

// Format2Bs 格式化文件大小速率，支持从 KB/s 到 EB/s 的单位转换
func Format2Bs(fileSize int64) string {
	if fileSize < 0 {
		return "——"
	}
	for _, u := range rateUnits {
		if float64(fileSize) < u.value {
			return fmt.Sprintf("%.2f %s", float64(fileSize)/u.value, u.unit)
		}
	}
	return fmt.Sprintf("%.2f EB/s", float64(fileSize)/float64(EB))
}

// FormatSize 调用 Format2Bs，格式化文件大小速率
func FormatSize(fileSize int64) string {
	return Format2Bs(fileSize)
}

// FormatSpeed 调用 Format2bps，格式化比特率
func FormatSpeed(speed float64) string {
	return Format2bps(speed)
}