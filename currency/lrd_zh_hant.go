//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lrd.RegisterName(xlanguage.MustParse("zh-Hant"), "賴比瑞亞元")
}
