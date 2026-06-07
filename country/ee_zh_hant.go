//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.MustParse("zh-Hant"), "愛沙尼亞")
	dataEstonia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "愛沙尼亞共和國")
	dataEstonia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "塔林")
}
