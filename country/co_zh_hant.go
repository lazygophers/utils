//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.MustParse("zh-Hant"), "哥倫比亞")
	dataColombia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "哥倫比亞共和國")
	dataColombia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "波哥大")
}
