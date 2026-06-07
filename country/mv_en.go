//go:build country_all || country_asia || country_mv || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.English, "Maldives")
	dataMaldives.RegisterOfficialName(xlanguage.English, "Republic of Maldives")
	dataMaldives.RegisterCapital(xlanguage.English, "Male")
}
