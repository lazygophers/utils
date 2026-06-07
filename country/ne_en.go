//go:build country_africa || country_all || country_ne || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.English, "Niger")
	dataNiger.RegisterOfficialName(xlanguage.English, "Republic of the Niger")
	dataNiger.RegisterCapital(xlanguage.English, "Niamey")
}
