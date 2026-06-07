//go:build lang_ar || lang_all

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	Yer.RegisterName(xlanguage.Arabic, "ريال يمني")
}
