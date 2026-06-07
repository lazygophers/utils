//go:build (lang_ko || lang_all) && (country_all || country_be || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Korean, "벨기에")
	dataBelgium.RegisterOfficialName(xlanguage.Korean, "벨기에 왕국")
	dataBelgium.RegisterCapital(xlanguage.Korean, "브뤼셀")
}
