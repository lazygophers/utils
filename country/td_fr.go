//go:build country_africa || country_all || country_middle_africa || country_td

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChad.RegisterName(xlanguage.French, "Tchad")
	dataChad.RegisterOfficialName(xlanguage.French, "République du Tchad")
	dataChad.RegisterCapital(xlanguage.French, "N'Djamena")
}
