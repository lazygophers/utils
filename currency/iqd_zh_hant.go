//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Iqd.RegisterName(xlanguage.MustParse("zh-Hant"), "伊拉克第納爾")
}
