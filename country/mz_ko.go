//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_mz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Korean, "모잠비크")
	dataMozambique.RegisterOfficialName(xlanguage.Korean, "모잠비크 공화국")
	dataMozambique.RegisterCapital(xlanguage.Korean, "마푸투")
}
