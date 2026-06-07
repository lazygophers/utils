//go:build country_africa || country_all || country_eastern_africa || country_km

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.French, "Comores")
	dataComoros.RegisterOfficialName(xlanguage.French, "Union des Comores")
	dataComoros.RegisterCapital(xlanguage.French, "Moroni")
}
