//go:build country_all || country_americas || country_caribbean || country_jm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.English, "Jamaica")
	dataJamaica.RegisterOfficialName(xlanguage.English, "Jamaica")
	dataJamaica.RegisterCapital(xlanguage.English, "Kingston")
}
