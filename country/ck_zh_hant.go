//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.MustParse("zh-Hant"), "庫克群島")
	dataCookIslands.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "庫克群島")
	dataCookIslands.RegisterCapital(xlanguage.MustParse("zh-Hant"), "阿瓦魯阿")
}
