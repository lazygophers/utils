//go:build (lang_ko || lang_all) && (country_all || country_asia || country_il || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.Korean, "이스라엘")
	dataIsrael.RegisterOfficialName(xlanguage.Korean, "이스라엘국")
	dataIsrael.RegisterCapital(xlanguage.Korean, "예루살렘")
}
