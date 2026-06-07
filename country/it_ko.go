//go:build (lang_ko || lang_all) && (country_all || country_europe || country_it || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataItaly.RegisterName(xlanguage.Korean, "이탈리아")
	dataItaly.RegisterOfficialName(xlanguage.Korean, "이탈리아 공화국")
	dataItaly.RegisterCapital(xlanguage.Korean, "로마")
}
