//go:build country_all || country_at || country_europe || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.English, "Austria")
	dataAustria.RegisterOfficialName(xlanguage.English, "Republic of Austria")
	dataAustria.RegisterCapital(xlanguage.English, "Vienna")
}
