//go:build (lang_ko || lang_all) && (country_all || country_by || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.Korean, "벨라루스")
	dataBelarus.RegisterOfficialName(xlanguage.Korean, "벨라루스 공화국")
	dataBelarus.RegisterCapital(xlanguage.Korean, "민스크")
}
