//go:build country_africa || country_all || country_bw || country_southern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.English, "Botswana")
	dataBotswana.RegisterOfficialName(xlanguage.English, "Republic of Botswana")
	dataBotswana.RegisterCapital(xlanguage.English, "Gaborone")
}
