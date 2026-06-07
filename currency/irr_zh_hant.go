//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Irr.RegisterName(xlanguage.MustParse("zh-Hant"), "伊朗里亞爾")
}
