//go:build (lang_es || lang_all) && (country_all || country_asia || country_ir || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIran.RegisterName(xlanguage.Spanish, "Irán")
	dataIran.RegisterOfficialName(xlanguage.Spanish, "República Islámica de Irán")
	dataIran.RegisterCapital(xlanguage.Spanish, "Teherán")
}
