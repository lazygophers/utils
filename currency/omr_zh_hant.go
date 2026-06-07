//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Omr.RegisterName(xlanguage.MustParse("zh-Hant"), "阿曼里亞爾")
}
