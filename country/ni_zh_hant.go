//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNicaragua.RegisterName(xlanguage.MustParse("zh-Hant"), "尼加拉瓜")
	dataNicaragua.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "尼加拉瓜共和國")
	dataNicaragua.RegisterCapital(xlanguage.MustParse("zh-Hant"), "馬拿瓜")
}
