//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Mzn.RegisterName(xlanguage.MustParse("zh-Hant"), "莫三比克梅蒂卡爾")
}
