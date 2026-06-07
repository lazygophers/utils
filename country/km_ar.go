//go:build country_africa || country_all || country_eastern_africa || country_km

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataComoros.RegisterName(xlanguage.Arabic, "جزر القمر")
	dataComoros.RegisterOfficialName(xlanguage.Arabic, "اتحاد جزر القمر")
	dataComoros.RegisterCapital(xlanguage.Arabic, "موروني")
}
