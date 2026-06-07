//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedStates.RegisterName(xlanguage.Japanese, "アメリカ合衆国")
	dataUnitedStates.RegisterOfficialName(xlanguage.Japanese, "アメリカ合衆国")
	dataUnitedStates.RegisterCapital(xlanguage.Japanese, "ワシントンD.C.")
}
