//go:build (lang_ko || lang_all) && (country_all || country_as || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.Korean, "아메리칸사모아")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.Korean, "아메리칸사모아")
	dataAmericanSamoa.RegisterCapital(xlanguage.Korean, "파고파고")
}
