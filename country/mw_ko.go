//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_mw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Korean, "말라위")
	dataMalawi.RegisterOfficialName(xlanguage.Korean, "말라위 공화국")
	dataMalawi.RegisterCapital(xlanguage.Korean, "릴롱궤")
}
