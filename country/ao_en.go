//go:build country_africa || country_all || country_ao || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.English, "Angola")
	dataAngola.RegisterOfficialName(xlanguage.English, "Republic of Angola")
	dataAngola.RegisterCapital(xlanguage.English, "Luanda")
}
