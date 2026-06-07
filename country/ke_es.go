//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_ke)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.Spanish, "Kenia")
	dataKenya.RegisterOfficialName(xlanguage.Spanish, "República de Kenia")
	dataKenya.RegisterCapital(xlanguage.Spanish, "Nairobi")
}
