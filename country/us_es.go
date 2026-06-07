//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.Spanish, "Estados Unidos")
	dataUnitedStates.RegisterOfficialName(xlanguage.Spanish, "Estados Unidos de América")
	dataUnitedStates.RegisterCapital(xlanguage.Spanish, "Washington D. C.")
}
