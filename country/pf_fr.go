//go:build country_all || country_oceania || country_pf || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchPolynesia.RegisterName(xlanguage.French, "Polynésie française")
	dataFrenchPolynesia.RegisterOfficialName(xlanguage.French, "Polynésie française")
	dataFrenchPolynesia.RegisterCapital(xlanguage.French, "Papeete")
}
