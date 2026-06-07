//go:build country_africa || country_all || country_ng || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.English, "Nigeria")
	dataNigeria.RegisterOfficialName(xlanguage.English, "Federal Republic of Nigeria")
	dataNigeria.RegisterCapital(xlanguage.English, "Abuja")
}
