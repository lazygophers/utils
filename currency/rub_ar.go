//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Rub.RegisterName(xlanguage.Arabic, "روبل روسي")
}
