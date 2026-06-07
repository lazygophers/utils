//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHungary.RegisterName(xlanguage.Spanish, "Hungría")
	dataHungary.RegisterOfficialName(xlanguage.Spanish, "Hungría")
	dataHungary.RegisterCapital(xlanguage.Spanish, "Budapest")
}
