//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Zmw.RegisterName(xlanguage.MustParse("zh-Hant"), "尚比亞克瓦查")
}
