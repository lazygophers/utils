//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.MustParse("zh-Hant"), "阿爾及利亞")
	dataAlgeria.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "阿爾及利亞人民民主共和國")
	dataAlgeria.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿爾及爾")
}
