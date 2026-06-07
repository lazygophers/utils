//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Gel.RegisterName(xlanguage.MustParse("zh-Hant"), "喬治亞拉里")
}
