//go:build country_ad || country_all || country_europe || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.English, "Andorra")
	dataAndorra.RegisterOfficialName(xlanguage.English, "Principality of Andorra")
	dataAndorra.RegisterCapital(xlanguage.English, "Andorra la Vella")
}
