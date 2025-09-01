package bufiox

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestDropCR 测试 dropCR 函数
func TestDropCR(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "空切片",
			input:    []byte{},
			expected: []byte{},
		},
		{
			name:     "无CR字符",
			input:    []byte("hello world"),
			expected: []byte("hello world"),
		},
		{
			name:     "以CR结尾",
			input:    []byte("hello world\r"),
			expected: []byte("hello world"),
		},
		{
			name:     "中间有CR但不在结尾",
			input:    []byte("hello\rworld"),
			expected: []byte("hello\rworld"),
		},
		{
			name:     "只有CR",
			input:    []byte("\r"),
			expected: []byte{},
		},
		{
			name:     "多个CR但只在结尾",
			input:    []byte("hello\r\nworld\r"),
			expected: []byte("hello\r\nworld"),
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			result := dropCR(tt.input)
			assert.Equal(t, tt.expected, result, "dropCR() 的结果应与期望值相等")
		})
	}
}

// TestScanBy 测试 ScanBy 函数
func TestScanBy(t *testing.T) {
	// 测试用逗号作为分隔符
	splitFunc := ScanBy([]byte(","))

	tests := []struct {
		name     string
		data     []byte
		atEOF    bool
		advance  int
		token    []byte
		err      error
	}{
		{
			name:     "空数据且EOF",
			data:     []byte{},
			atEOF:    true,
			advance:  0,
			token:    nil,
			err:      nil,
		},
		{
			name:     "找到分隔符",
			data:     []byte("hello,world"),
			atEOF:    false,
			advance:  6, // "hello," 的长度
			token:    []byte("hello"),
			err:      nil,
		},
		{
			name:     "未找到分隔符且未EOF",
			data:     []byte("hello"),
			atEOF:    false,
			advance:  0,
			token:    nil,
			err:      nil,
		},
		{
			name:     "未找到分隔符但已EOF",
			data:     []byte("hello"),
			atEOF:    true,
			advance:  5, // "hello" 的长度
			token:    []byte("hello"),
			err:      nil,
		},
		{
			name:     "多个分隔符，取第一个",
			data:     []byte("a,b,c"),
			atEOF:    false,
			advance:  2, // "a," 的长度
			token:    []byte("a"),
			err:      nil,
		},
		{
			name:     "分隔符在开头",
			data:     []byte(",start"),
			atEOF:    false,
			advance:  1, // "," 的长度
			token:    []byte{},
			err:      nil,
		},
		{
			name:     "分隔符在结尾且EOF",
			data:     []byte("end,"),
			atEOF:    true,
			advance:  4, // "end," 的长度
			token:    []byte("end"),
			err:      nil,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			advance, token, err := splitFunc(tt.data, tt.atEOF)
			assert.Equal(t, tt.advance, advance, "advance 应与期望值相等")
			assert.Equal(t, tt.token, token, "token 应与期望值相等")
			assert.Equal(t, tt.err, err, "err 应与期望值相等")
		})
	}
}

// TestScanByWithEmptySeparator 测试使用空分隔符的情况
func TestScanByWithEmptySeparator(t *testing.T) {
	splitFunc := ScanBy([]byte{})

	data := []byte("hello world")
	advance, token, err := splitFunc(data, false)
	
	// 空分隔符返回0 advance和空token
	assert.Equal(t, 0, advance, "使用空分隔符时应返回0 advance")
	assert.Equal(t, []byte{}, token, "使用空分隔符时应返回空token")
	assert.Nil(t, err, "err 应为 nil")
}

// TestScanLines 测试 ScanLines 函数
func TestScanLines(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		atEOF    bool
		advance  int
		token    []byte
		err      error
	}{
		{
			name:     "空数据且EOF",
			data:     []byte{},
			atEOF:    true,
			advance:  0,
			token:    nil,
			err:      nil,
		},
		{
			name:     "Unix换行符",
			data:     []byte("hello\nworld"),
			atEOF:    false,
			advance:  6, // "hello\n" 的长度
			token:    []byte("hello"),
			err:      nil,
		},
		{
			name:     "Windows换行符",
			data:     []byte("hello\r\nworld"),
			atEOF:    false,
			advance:  7, // "hello\r\n" 的长度是7
			token:    []byte("hello"),
			err:      nil,
		},
		{
			name:     "只有CR",
			data:     []byte("hello\rworld"),
			atEOF:    false,
			advance:  0, // ScanLines 不处理单独的CR作为换行符，只处理LF
			token:    nil,
			err:      nil,
		},
		{
			name:     "未找到换行符且未EOF",
			data:     []byte("hello"),
			atEOF:    false,
			advance:  0,
			token:    nil,
			err:      nil,
		},
		{
			name:     "未找到换行符但已EOF",
			data:     []byte("hello"),
			atEOF:    true,
			advance:  5, // "hello" 的长度
			token:    []byte("hello"),
			err:      nil,
		},
		{
			name:     "只有换行符",
			data:     []byte("\n"),
			atEOF:    false,
			advance:  1,
			token:    []byte{},
			err:      nil,
		},
		{
			name:     "只有Windows换行符",
			data:     []byte("\r\n"),
			atEOF:    false,
			advance:  2,
			token:    []byte{},
			err:      nil,
		},
		{
			name:     "多行数据",
			data:     []byte("line1\nline2\nline3"),
			atEOF:    false,
			advance:  6,
			token:    []byte("line1"),
			err:      nil,
		},
	}

	for _, tt := range tests {
		tt := tt // 避免竞态
		t.Run(tt.name, func(t *testing.T) {
			advance, token, err := ScanLines(tt.data, tt.atEOF)
			assert.Equal(t, tt.advance, advance, "advance 应与期望值相等")
			assert.Equal(t, tt.token, token, "token 应与期望值相等")
			assert.Equal(t, tt.err, err, "err 应与期望值相等")
		})
	}
}

// TestScanLinesIntegration 测试 ScanLines 与 bufio.Scanner 的集成
func TestScanLinesIntegration(t *testing.T) {
	data := "line1\nline2\r\nline3\nline4"
	scanner := bufio.NewScanner(bytes.NewBufferString(data))
	scanner.Split(ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	require.NoError(t, scanner.Err(), "Scanner 不应有错误")
	assert.Equal(t, []string{"line1", "line2", "line3", "line4"}, lines, "扫描的行应与期望值相等")
}

// TestScanLinesWithFinalLineWithoutNewline 测试最后一行没有换行符的情况
func TestScanLinesWithFinalLineWithoutNewline(t *testing.T) {
	data := "line1\nline2"
	scanner := bufio.NewScanner(bytes.NewBufferString(data))
	scanner.Split(ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	require.NoError(t, scanner.Err(), "Scanner 不应有错误")
	assert.Equal(t, []string{"line1", "line2"}, lines, "扫描的行应与期望值相等")
}

// TestScanByWithMultibyteSeparator 测试多字节分隔符
func TestScanByWithMultibyteSeparator(t *testing.T) {
	separator := []byte("||")
	splitFunc := ScanBy(separator)

	data := []byte("part1||part2||part3")
	advance, token, err := splitFunc(data, false)

	assert.Equal(t, 7, advance, "advance 应为 7 (part1||)")
	assert.Equal(t, []byte("part1"), token, "token 应为 part1")
	assert.Nil(t, err, "err 应为 nil")
}

// TestDropCRPerformance 测试 dropCR 的性能
func TestDropCRPerformance(t *testing.T) {
	// 构造一个大切片
	largeData := make([]byte, 1024*1024) // 1MB
	for i := 0; i < len(largeData)-1; i++ {
		largeData[i] = 'a'
	}
	largeData[len(largeData)-1] = '\r'

	t.Run("LargeDataWithCR", func(t *testing.T) {
		result := dropCR(largeData)
		assert.Equal(t, len(largeData)-1, len(result), "结果长度应比原数据少1")
		assert.Equal(t, byte('a'), result[len(result)-1], "最后一个字符应为 'a'")
	})
}