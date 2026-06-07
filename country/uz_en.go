//go:build country_all || country_asia || country_central_asia || country_uz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUzbekistan.RegisterName(xlanguage.English, "Uzbekistan")
	dataUzbekistan.RegisterOfficialName(xlanguage.English, "Republic of Uzbekistan")
	dataUzbekistan.RegisterCapital(xlanguage.English, "Tashkent")
}
