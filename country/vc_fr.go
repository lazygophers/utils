//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.French, "Saint-Vincent-et-les Grenadines")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.French, "Saint-Vincent-et-les Grenadines")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.French, "Kingstown")
}
