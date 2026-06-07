//go:build (lang_fr || lang_all) && (country_all || country_asia || country_ir || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.French, "Iran")
	dataIran.RegisterOfficialName(xlanguage.French, "République islamique d'Iran")
	dataIran.RegisterCapital(xlanguage.French, "Téhéran")
}
