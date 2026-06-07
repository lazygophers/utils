//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Azn.RegisterName(xlanguage.MustParse("zh-Hant"), "亞塞拜然馬納特")
}
