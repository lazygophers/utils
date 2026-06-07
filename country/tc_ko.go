//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_tc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurksAndCaicosIslands.RegisterName(xlanguage.Korean, "터크스 케이커스 제도")
	dataTurksAndCaicosIslands.RegisterOfficialName(xlanguage.Korean, "터크스 케이커스 제도")
	dataTurksAndCaicosIslands.RegisterCapital(xlanguage.Korean, "코크번타운")
}
