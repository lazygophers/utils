//go:build (lang_fr || lang_all) && (country_ag || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.French, "Antigua-et-Barbuda")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.French, "Antigua-et-Barbuda")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.French, "Saint John's")
}
