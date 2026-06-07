//go:build country_africa || country_all || country_bi || country_eastern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.English, "Burundi")
	dataBurundi.RegisterOfficialName(xlanguage.English, "Republic of Burundi")
	dataBurundi.RegisterCapital(xlanguage.English, "Gitega")
}
