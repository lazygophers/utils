//go:build country_africa || country_all || country_gn || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuinea.RegisterName(xlanguage.English, "Guinea")
	dataGuinea.RegisterOfficialName(xlanguage.English, "Republic of Guinea")
	dataGuinea.RegisterCapital(xlanguage.English, "Conakry")
}
