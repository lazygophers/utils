//go:build (lang_zh_hant || lang_all) && (country_all || country_asia || country_eastern_asia || country_mn || currency_all || currency_mnt)

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MNT.RegisterName(xlanguage.MustParse("zh-Hant"), "圖格里克")
}
