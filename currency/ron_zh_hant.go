//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ron.RegisterName(xlanguage.MustParse("zh-Hant"), "羅馬尼亞列伊")
}
