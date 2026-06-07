//go:build (lang_ko || lang_all) && (country_all || country_ee || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Korean, "에스토니아")
	dataEstonia.RegisterOfficialName(xlanguage.Korean, "에스토니아 공화국")
	dataEstonia.RegisterCapital(xlanguage.Korean, "탈린")
}
