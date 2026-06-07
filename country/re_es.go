//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_re)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Spanish, "Reunión")
	dataReunion.RegisterOfficialName(xlanguage.Spanish, "Reunión")
	dataReunion.RegisterCapital(xlanguage.Spanish, "Saint-Denis")
}
