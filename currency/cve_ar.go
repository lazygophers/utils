//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Cve.RegisterName(xlanguage.Arabic, "إسكودو الرأس الأخضر")
}
