//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.MustParse("zh-Hant"), "奈及利亞")
	dataNigeria.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "奈及利亞聯邦共和國")
	dataNigeria.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿布加")
}
