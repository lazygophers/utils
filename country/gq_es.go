//go:build country_africa || country_all || country_gq || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEquatorialGuinea.RegisterName(xlanguage.Spanish, "Guinea Ecuatorial")
	dataEquatorialGuinea.RegisterOfficialName(xlanguage.Spanish, "República de Guinea Ecuatorial")
	dataEquatorialGuinea.RegisterCapital(xlanguage.Spanish, "Malabo")
}
