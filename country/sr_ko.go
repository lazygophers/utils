//go:build (lang_ko || lang_all) && (country_all || country_americas || country_south_america || country_sr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Korean, "수리남")
	dataSuriname.RegisterOfficialName(xlanguage.Korean, "수리남 공화국")
	dataSuriname.RegisterCapital(xlanguage.Korean, "파라마리보")
}
