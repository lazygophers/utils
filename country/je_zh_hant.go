//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.MustParse("zh-Hant"), "澤西")
	dataJersey.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "澤西行政區")
	dataJersey.RegisterCapital(xlanguage.MustParse("zh-Hant"), "聖赫利爾")
}
