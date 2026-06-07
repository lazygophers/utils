//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Dzd.RegisterName(xlanguage.MustParse("zh-Hant"), "阿爾及利亞第納爾")
}
