//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Arabic, "المجر")
	dataHungary.RegisterOfficialName(xlanguage.Arabic, "المجر")
	dataHungary.RegisterCapital(xlanguage.Arabic, "بودابست")
}
