//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_re)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Korean, "레위니옹")
	dataReunion.RegisterOfficialName(xlanguage.Korean, "레위니옹")
	dataReunion.RegisterCapital(xlanguage.Korean, "생드니")
}
