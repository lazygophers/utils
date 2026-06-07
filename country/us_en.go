package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.English, "United States")
	dataUnitedStates.RegisterOfficialName(xlanguage.English, "United States of America")
	dataUnitedStates.RegisterCapital(xlanguage.English, "Washington, D.C.")
}
