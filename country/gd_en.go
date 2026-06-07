//go:build country_all || country_americas || country_caribbean || country_gd

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.English, "Grenada")
	dataGrenada.RegisterOfficialName(xlanguage.English, "Grenada")
	dataGrenada.RegisterCapital(xlanguage.English, "Saint George's")
}
