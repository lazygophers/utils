//go:build country_africa || country_all || country_eastern_africa || country_mg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.English, "Madagascar")
	dataMadagascar.RegisterOfficialName(xlanguage.English, "Republic of Madagascar")
	dataMadagascar.RegisterCapital(xlanguage.English, "Antananarivo")
}
