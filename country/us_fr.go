//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.French, "États-Unis")
	dataUnitedStates.RegisterOfficialName(xlanguage.French, "États-Unis d'Amérique")
	dataUnitedStates.RegisterCapital(xlanguage.French, "Washington")
}
