//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.MustParse("zh-Hant"), "印度")
	dataIndia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "印度共和國")
	dataIndia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "新德里")
}
