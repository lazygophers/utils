//go:build (lang_fr || lang_all) && (country_all || country_europe || country_northern_europe || country_sj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.French, "Svalbard et Jan Mayen")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.French, "Svalbard et Jan Mayen")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.French, "Longyearbyen")
}
