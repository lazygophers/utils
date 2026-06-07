//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJapan.RegisterName(xlanguage.MustParse("zh-Hant"), "日本")
	dataJapan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "日本")
	dataJapan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "東京")
}
