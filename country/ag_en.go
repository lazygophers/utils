//go:build country_ag || country_all || country_americas || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.English, "Antigua and Barbuda")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.English, "Antigua and Barbuda")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.English, "Saint John's")
}
