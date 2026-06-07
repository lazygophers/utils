//go:build (lang_ko || lang_all) && (country_all || country_asia || country_western_asia || country_ye)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataYemen.RegisterName(xlanguage.Korean, "예멘")
	dataYemen.RegisterOfficialName(xlanguage.Korean, "예멘 공화국")
	dataYemen.RegisterCapital(xlanguage.Korean, "사나")
}
