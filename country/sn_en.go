//go:build country_africa || country_all || country_sn || country_western_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.English, "Senegal")
	dataSenegal.RegisterOfficialName(xlanguage.English, "Republic of Senegal")
	dataSenegal.RegisterCapital(xlanguage.English, "Dakar")
}
