//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Lak.RegisterName(xlanguage.Arabic, "كيب لاوسي")
}
