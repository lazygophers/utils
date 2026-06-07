//go:build (lang_ko || lang_all) && (country_all || country_oceania || country_polynesia || country_ws)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSamoa.RegisterName(xlanguage.Korean, "사모아")
	dataSamoa.RegisterOfficialName(xlanguage.Korean, "사모아 독립국")
	dataSamoa.RegisterCapital(xlanguage.Korean, "아피아")
}
