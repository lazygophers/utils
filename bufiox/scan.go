package bufiox

import "bytes"

// ScanBy 创建并返回一个自定义的扫描函数，用于按指定字节序列分割数据
// 参数 seq: 需要搜索的字节分隔符（如 []byte("\n") 表示换行符分割）
// 返回值: 返回的扫描函数接收字节数据和EOF标志，返回三个参数：
//
//	advance: 指示应前进的字节数
//	token: 当前分割出的有效数据标记
//	err: 错误信息（若返回nil表示继续处理）
//
// 函数特性：
//  1. 当atEOF为true且无数据时立即返回nil
//  2. 支持非EOF状态下的分块处理
//  3. 返回的分割函数遵循bufio.Scanner接口规范
//
// ScanBy 返回一个扫描函数，该函数根据指定的字节序列分割输入
// 参数 seq 是分隔符字节序列
// 返回的函数符合 bufio.SplitFunc 接口规范
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

// ScanLines 实现基于bufio.Scanner的行扫描器，处理跨平台换行符分割
// 参数 data: 当前处理的字节切片
// 参数 atEOF: 是否为数据流结束标志
// 返回值: 三元组(advance, token, err)
// 函数特性：
//  1. 自动处理CRLF（Windows）和LF（Unix）换行符
//  2. 在EOF时强制分割剩余数据
//  3. 通过dropCR函数去除Windows换行符中的CR
//  4. 遵循标准bufio.Scanner接口规范
//
// ScanLines 实现按行分割的扫描函数
// 处理换行符(\n)分割逻辑，特别处理Windows换行符(\r\n)场景
// 参数:
//   - data: 当前处理的字节数据
//   - atEOF: 是否已到达数据末尾
//
// 返回:
//   - advance: 消耗的字节数
//   - token: 当前分割得到的行数据（已处理CRLF）
//   - err: 错误信息
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

// dropCR 是行处理辅助函数，用于规范换行符格式
// 功能：移除字节序列最后一个字符的CR（\r）
// 使用场景：当处理Windows风格换行符CRLF（\r\n）时
// 返回值: 返回处理后的字节切片（若存在CR则去除，否则原样返回）
// dropCR 用于移除字节切片末尾的回车符(\r)
// 主要处理Windows换行符场景，将CRLF转换为LF
// 参数 data: 需要处理的字节切片
// 返回: 移除末尾\r后的字节切片
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
