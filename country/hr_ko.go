//go:build (lang_ko || lang_all) && (country_all || country_europe || country_hr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.Korean, "크로아티아")
	dataCroatia.RegisterOfficialName(xlanguage.Korean, "크로아티아 공화국")
	dataCroatia.RegisterCapital(xlanguage.Korean, "자그레브")
}
