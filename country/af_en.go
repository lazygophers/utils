//go:build country_af || country_all || country_asia || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.English, "Afghanistan")
	dataAfghanistan.RegisterOfficialName(xlanguage.English, "Islamic Emirate of Afghanistan")
	dataAfghanistan.RegisterCapital(xlanguage.English, "Kabul")
}
