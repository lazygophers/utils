//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dkk.RegisterName(xlanguage.MustParse("zh-Hant"), "丹麥克朗")
}
