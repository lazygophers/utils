//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.MustParse("zh-Hant"), "愛爾蘭")
	dataIreland.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "愛爾蘭共和國")
	dataIreland.RegisterCapital(xlanguage.MustParse("zh-Hant"), "都柏林")
}
