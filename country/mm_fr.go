//go:build (lang_fr || lang_all) && (country_all || country_asia || country_mm || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.French, "Birmanie")
	dataMyanmar.RegisterOfficialName(xlanguage.French, "République de l'Union du Myanmar")
	dataMyanmar.RegisterCapital(xlanguage.French, "Naypyidaw")
}
