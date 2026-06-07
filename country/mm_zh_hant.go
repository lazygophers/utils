//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.MustParse("zh-Hant"), "緬甸")
	dataMyanmar.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "緬甸聯邦共和國")
	dataMyanmar.RegisterCapital(xlanguage.MustParse("zh-Hant"), "奈比多")
}
