package bufiox

import "bytes"

// ScanBy 创建并返回一个自定义的扫描函数，用于按指定字节序列分割数据
func ScanBy(seq []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, seq); i >= 0 {
			// We have a full newline-terminated line.
			return i + len(seq), data[0:i], nil
		}

		if atEOF {
			return len(data), data, nil
		}

		return 0, nil, err
	}
}

// ScanLines 实现按行分割的扫描函数，兼容 CRLF 和 LF 换行符
// 与标准库 bufio.ScanLines 实现一致，自动处理 Windows 换行符场景
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR 移除字节切片末尾的回车符（\r），用于将 CRLF 转换为 LF
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
