//go:build country_africa || country_all || country_dj || country_eastern_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.English, "Djibouti")
	dataDjibouti.RegisterOfficialName(xlanguage.English, "Republic of Djibouti")
	dataDjibouti.RegisterCapital(xlanguage.English, "Djibouti")
}
