//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.French, "Hongrie")
	dataHungary.RegisterOfficialName(xlanguage.French, "Hongrie")
	dataHungary.RegisterCapital(xlanguage.French, "Budapest")
}
