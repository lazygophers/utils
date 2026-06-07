//go:build (lang_fr || lang_all) && (country_all || country_asia || country_central_asia || country_tm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.French, "Turkménistan")
	dataTurkmenistan.RegisterOfficialName(xlanguage.French, "Turkménistan")
	dataTurkmenistan.RegisterCapital(xlanguage.French, "Achgabat")
}
