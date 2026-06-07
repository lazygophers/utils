//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Sar.RegisterName(xlanguage.MustParse("zh-Hant"), "沙烏地里亞爾")
}
