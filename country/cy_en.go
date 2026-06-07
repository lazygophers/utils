//go:build country_all || country_cy || country_europe || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.English, "Cyprus")
	dataCyprus.RegisterOfficialName(xlanguage.English, "Republic of Cyprus")
	dataCyprus.RegisterCapital(xlanguage.English, "Nicosia")
}
