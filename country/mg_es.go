//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_mg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.Spanish, "Madagascar")
	dataMadagascar.RegisterOfficialName(xlanguage.Spanish, "República de Madagascar")
	dataMadagascar.RegisterCapital(xlanguage.Spanish, "Antananarivo")
}
