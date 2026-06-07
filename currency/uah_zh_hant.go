//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Uah.RegisterName(xlanguage.MustParse("zh-Hant"), "烏克蘭格里夫納")
}
