//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.MustParse("zh-Hant"), "塔吉克")
	dataTajikistan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "塔吉克共和國")
	dataTajikistan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "杜尚別")
}
