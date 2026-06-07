//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cve.RegisterName(xlanguage.MustParse("zh-Hant"), "維德角埃斯庫多")
}
