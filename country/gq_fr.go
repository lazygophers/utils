//go:build country_africa || country_all || country_gq || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.French, "Guinée équatoriale")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.French, "République de Guinée équatoriale")
	dataEquatorialGuinea.RegisterCapital(xlanguage.French, "Malabo")
}
