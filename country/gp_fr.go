//go:build country_all || country_americas || country_caribbean || country_gp

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuadeloupe.RegisterName(xlanguage.French, "Guadeloupe")
	dataGuadeloupe.RegisterOfficialName(xlanguage.French, "Guadeloupe")
	dataGuadeloupe.RegisterCapital(xlanguage.French, "Basse-Terre")
}
