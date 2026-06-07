//go:build country_all || country_oceania || country_pf || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.English, "French Polynesia")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.English, "French Polynesia")
	dataFrenchPolynesia.RegisterCapital(xlanguage.English, "Papeete")
}
