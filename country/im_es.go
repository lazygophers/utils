//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsleOfMan.RegisterName(xlanguage.Spanish, "Isla de Man")
	dataIsleOfMan.RegisterOfficialName(xlanguage.Spanish, "Isla de Man")
	dataIsleOfMan.RegisterCapital(xlanguage.Spanish, "Douglas")
}
