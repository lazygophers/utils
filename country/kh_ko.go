//go:build (lang_ko || lang_all) && (country_all || country_asia || country_kh || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCambodia.RegisterName(xlanguage.Korean, "캄보디아")
	dataCambodia.RegisterOfficialName(xlanguage.Korean, "캄보디아 왕국")
	dataCambodia.RegisterCapital(xlanguage.Korean, "프놈펜")
}
