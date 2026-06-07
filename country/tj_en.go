//go:build country_all || country_asia || country_central_asia || country_tj

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.English, "Tajikistan")
	dataTajikistan.RegisterOfficialName(xlanguage.English, "Republic of Tajikistan")
	dataTajikistan.RegisterCapital(xlanguage.English, "Dushanbe")
}
