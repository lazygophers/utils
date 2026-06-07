//go:build country_africa || country_all || country_cd || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.French, "République démocratique du Congo")
	dataDrCongo.RegisterOfficialName(xlanguage.French, "République démocratique du Congo")
	dataDrCongo.RegisterCapital(xlanguage.French, "Kinshasa")
}
