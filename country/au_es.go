//go:build (lang_es || lang_all) && (country_all || country_au || country_australia_and_new_zealand || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.Spanish, "Australia")
	dataAustralia.RegisterOfficialName(xlanguage.Spanish, "Mancomunidad de Australia")
	dataAustralia.RegisterCapital(xlanguage.Spanish, "Canberra")
}
