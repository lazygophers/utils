//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLibya.RegisterName(xlanguage.MustParse("zh-Hant"), "利比亞")
	dataLibya.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "利比亞國")
	dataLibya.RegisterCapital(xlanguage.MustParse("zh-Hant"), "的黎波里")
}
