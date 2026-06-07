//go:build country_all || country_europe || country_me || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.English, "Montenegro")
	dataMontenegro.RegisterOfficialName(xlanguage.English, "Montenegro")
	dataMontenegro.RegisterCapital(xlanguage.English, "Podgorica")
}
