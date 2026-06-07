//go:build country_africa || country_all || country_lr || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.English, "Liberia")
	dataLiberia.RegisterOfficialName(xlanguage.English, "Republic of Liberia")
	dataLiberia.RegisterCapital(xlanguage.English, "Monrovia")
}
