//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLithuania.RegisterName(xlanguage.MustParse("zh-Hant"), "立陶宛")
	dataLithuania.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "立陶宛共和國")
	dataLithuania.RegisterCapital(xlanguage.MustParse("zh-Hant"), "維爾紐斯")
}
