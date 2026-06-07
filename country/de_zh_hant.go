//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.MustParse("zh-Hant"), "德國")
	dataGermany.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "德意志聯邦共和國")
	dataGermany.RegisterCapital(xlanguage.MustParse("zh-Hant"), "柏林")
}
