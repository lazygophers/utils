//go:build !fake_en && !fake_zh_cn && !fake_zh_tw && !fake_fr && !fake_ru && !fake_pt && !fake_es

package fake

// 默认构建标签：包含所有语言支持
// 当没有指定特定语言标签时，使用此文件

func init() {
	// 预加载所有支持的语言数据（可选）
	// 在实际使用中，数据会按需加载
}