//go:build (lang_ko || lang_all) && (country_all || country_europe || country_sm || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSanMarino.RegisterName(xlanguage.Korean, "산마리노")
	dataSanMarino.RegisterOfficialName(xlanguage.Korean, "산마리노 공화국")
	dataSanMarino.RegisterCapital(xlanguage.Korean, "산마리노")
}
