//go:build (lang_es || lang_all) && (country_ag || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.Spanish, "Antigua y Barbuda")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.Spanish, "Antigua y Barbuda")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.Spanish, "Saint John's")
}
