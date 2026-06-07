//go:build (lang_es || lang_all) && (country_africa || country_all || country_eastern_africa || country_km)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.Spanish, "Comoras")
	dataComoros.RegisterOfficialName(xlanguage.Spanish, "Unión de las Comoras")
	dataComoros.RegisterCapital(xlanguage.Spanish, "Moroni")
}
