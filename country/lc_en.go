//go:build country_all || country_americas || country_caribbean || country_lc

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintLucia.RegisterName(xlanguage.English, "Saint Lucia")
	dataSaintLucia.RegisterOfficialName(xlanguage.English, "Saint Lucia")
	dataSaintLucia.RegisterCapital(xlanguage.English, "Castries")
}
