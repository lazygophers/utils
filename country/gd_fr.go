//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_gd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGrenada.RegisterName(xlanguage.French, "Grenade")
	dataGrenada.RegisterOfficialName(xlanguage.French, "Grenade")
	dataGrenada.RegisterCapital(xlanguage.French, "Saint-Georges")
}
