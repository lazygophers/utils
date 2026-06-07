//go:build country_africa || country_all || country_eastern_africa || country_km

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.English, "Comoros")
	dataComoros.RegisterOfficialName(xlanguage.English, "Union of the Comoros")
	dataComoros.RegisterCapital(xlanguage.English, "Moroni")
}
