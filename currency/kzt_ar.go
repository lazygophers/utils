//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Kzt.RegisterName(xlanguage.Arabic, "تنغة كازاخستاني")
}
