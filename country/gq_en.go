//go:build country_africa || country_all || country_gq || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.English, "Equatorial Guinea")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.English, "Republic of Equatorial Guinea")
	dataEquatorialGuinea.RegisterCapital(xlanguage.English, "Malabo")
}
