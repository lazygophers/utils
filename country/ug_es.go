//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_ug)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUganda.RegisterName(xlanguage.Spanish, "Uganda")
	dataUganda.RegisterOfficialName(xlanguage.Spanish, "República de Uganda")
	dataUganda.RegisterCapital(xlanguage.Spanish, "Kampala")
}
