//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Wst.RegisterName(xlanguage.Arabic, "تالا ساموي")
}
