//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Spanish, "Seychelles")
	dataSeychelles.RegisterOfficialName(xlanguage.Spanish, "República de Seychelles")
	dataSeychelles.RegisterCapital(xlanguage.Spanish, "Victoria")
}
