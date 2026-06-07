//go:build lang_zh_hant || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Ils.RegisterName(xlanguage.MustParse("zh-Hant"), "新以色列謝克爾")
}
