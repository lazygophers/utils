//go:build country_all || country_asia || country_central_asia || country_tm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.English, "Turkmenistan")
	dataTurkmenistan.RegisterOfficialName(xlanguage.English, "Turkmenistan")
	dataTurkmenistan.RegisterCapital(xlanguage.English, "Ashgabat")
}
