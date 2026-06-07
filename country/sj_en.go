//go:build country_all || country_europe || country_northern_europe || country_sj

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.English, "Svalbard and Jan Mayen")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.English, "Svalbard and Jan Mayen")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.English, "Longyearbyen")
}
