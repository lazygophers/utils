//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Syp.RegisterName(xlanguage.MustParse("zh-Hant"), "敘利亞鎊")
}
