//go:build (lang_ko || lang_all) && (country_all || country_oceania || country_polynesia || country_tv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTuvalu.RegisterName(xlanguage.Korean, "투발루")
	dataTuvalu.RegisterOfficialName(xlanguage.Korean, "투발루")
	dataTuvalu.RegisterCapital(xlanguage.Korean, "푸나푸티")
}
