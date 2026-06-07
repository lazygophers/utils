//go:build country_all || country_americas || country_caribbean || country_gp

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.English, "Guadeloupe")
	dataGuadeloupe.RegisterOfficialName(xlanguage.English, "Guadeloupe")
	dataGuadeloupe.RegisterCapital(xlanguage.English, "Basse-Terre")
}
