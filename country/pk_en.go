//go:build country_all || country_asia || country_pk || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.English, "Pakistan")
	dataPakistan.RegisterOfficialName(xlanguage.English, "Islamic Republic of Pakistan")
	dataPakistan.RegisterCapital(xlanguage.English, "Islamabad")
}
