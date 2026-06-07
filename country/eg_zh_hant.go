//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEgypt.RegisterName(xlanguage.MustParse("zh-Hant"), "埃及")
	dataEgypt.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿拉伯埃及共和國")
	dataEgypt.RegisterCapital(xlanguage.MustParse("zh-Hant"), "開羅")
}
