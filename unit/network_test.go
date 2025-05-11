package unit

import (
	"fmt"
	"testing"
)

// TestFormat2bps 测试 Format2bps 函数
func TestFormat2bps(t *testing.T) {
	testCases := []struct {
		speed float64
		want  string
	}{
		{0, "——"},
		{1024, "1.00 Kbps"},
		{1024 * 1024, "1.00 Mbps"},
		{1024 * 1024 * 1024, "1.00 Gbps"},
		{1024 * 1024 * 1024 * 1024, "1.00 Tbps"},
		{1024 * 1024 * 1024 * 1024 * 1024, "1.00 Pbps"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1.00 Ebps"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("speed=%f", tc.speed), func(t *testing.T) {
			got := Format2bps(tc.speed)
			if got != tc.want {
				t.Errorf("Format2bps(%f) = %s; want %s", tc.speed, got, tc.want)
			}
		})
	}
}

// TestFormat2B 测试 Format2B 函数
func TestFormat2B(t *testing.T) {
	testCases := []struct {
		fileSize int64
		want     string
	}{
		{-1, "——"},
		{0, "0.00 KB"},
		{1024, "1.00 KB"},
		{1024 * 1024, "1.00 MB"},
		{1024 * 1024 * 1024, "1.00 GB"},
		{1024 * 1024 * 1024 * 1024, "1.00 TB"},
		{1024 * 1024 * 1024 * 1024 * 1024, "1.00 PB"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1.00 EB"}, // 确保 EB 单位的测试用例正确
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("fileSize=%d", tc.fileSize), func(t *testing.T) {
			got := Format2B(tc.fileSize)
			if got != tc.want {
				t.Errorf("Format2B(%d) = %s; want %s", tc.fileSize, got, tc.want)
			}
		})
	}
}

// TestFormat2Bs 测试 Format2Bs 函数
func TestFormat2Bs(t *testing.T) {
	testCases := []struct {
		fileSize int64
		want     string
	}{
		{-1, "——"},
		{0, "——"},
		{8 * 1024, "1.00 KB/s"},
		{8 * 1024 * 1024, "1.00 MB/s"},
		{8 * 1024 * 1024 * 1024, "1.00 GB/s"},
		{8 * 1024 * 1024 * 1024 * 1024, "1.00 TB/s"},
		{8 * 1024 * 1024 * 1024 * 1024 * 1024, "1.00 PB/s"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("fileSize=%d", tc.fileSize), func(t *testing.T) {
			got := Format2Bs(float64(tc.fileSize))
			if got != tc.want {
				t.Errorf("Format2Bs(%d) = %s; want %s", tc.fileSize, got, tc.want)
			}
		})
	}
}

// TestFormatSize 测试 FormatSize 函数
func TestFormatSize(t *testing.T) {
	testCases := []struct {
		fileSize int64
		want     string
	}{
		{-1, "——"},
		{0, "0.00 KB"},
		{1024, "1.00 KB"},
		{1024 * 1024, "1.00 MB"},
		{1024 * 1024 * 1024, "1.00 GB"},
		{1024 * 1024 * 1024 * 1024, "1.00 TB"},
		{1024 * 1024 * 1024 * 1024 * 1024, "1.00 PB"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1.00 EB"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("fileSize=%d", tc.fileSize), func(t *testing.T) {
			got := FormatSize(tc.fileSize)
			if got != tc.want {
				t.Errorf("FormatSize(%d) = %s; want %s", tc.fileSize, got, tc.want)
			}
		})
	}
}

// TestFormatSpeed 测试 FormatSpeed 函数
func TestFormatSpeed(t *testing.T) {
	testCases := []struct {
		speed float64
		want  string
	}{
		{0, "——"},
		{1024, "1.00 Kbps"},
		{1024 * 1024, "1.00 Mbps"},
		{1024 * 1024 * 1024, "1.00 Gbps"},
		{1024 * 1024 * 1024 * 1024, "1.00 Tbps"},
		{1024 * 1024 * 1024 * 1024 * 1024, "1.00 Pbps"},
		{1024 * 1024 * 1024 * 1024 * 1024 * 1024, "1.00 Ebps"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("speed=%f", tc.speed), func(t *testing.T) {
			got := FormatSpeed(tc.speed)
			if got != tc.want {
				t.Errorf("FormatSpeed(%f) = %s; want %s", tc.speed, got, tc.want)
			}
		})
	}
}
