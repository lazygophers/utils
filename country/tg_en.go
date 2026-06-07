//go:build country_africa || country_all || country_tg || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.English, "Togo")
	dataTogo.RegisterOfficialName(xlanguage.English, "Togolese Republic")
	dataTogo.RegisterCapital(xlanguage.English, "Lome")
}
