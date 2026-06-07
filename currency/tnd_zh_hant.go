//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Tnd.RegisterName(xlanguage.MustParse("zh-Hant"), "突尼西亞第納爾")
}
