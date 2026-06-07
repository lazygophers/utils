//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_mu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.Korean, "모리셔스")
	dataMauritius.RegisterOfficialName(xlanguage.Korean, "모리셔스 공화국")
	dataMauritius.RegisterCapital(xlanguage.Korean, "포트루이스")
}
