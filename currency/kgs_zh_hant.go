//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kgs.RegisterName(xlanguage.MustParse("zh-Hant"), "吉爾吉斯索姆")
}
