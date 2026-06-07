//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Bob.RegisterName(xlanguage.MustParse("zh-Hant"), "玻利維亞諾")
}
