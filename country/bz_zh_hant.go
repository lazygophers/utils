//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.MustParse("zh-Hant"), "貝里斯")
	dataBelize.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "貝里斯")
	dataBelize.RegisterCapital(xlanguage.MustParse("zh-Hant"), "貝爾墨邦")
}
