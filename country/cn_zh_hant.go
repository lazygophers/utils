//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.MustParse("zh-Hant"), "中國")
	dataChina.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "中華人民共和國")
	dataChina.RegisterCapital(xlanguage.MustParse("zh-Hant"), "北京")
}
