//go:build (lang_ko || lang_all) && (country_all || country_europe || country_is || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.Korean, "아이슬란드")
	dataIceland.RegisterOfficialName(xlanguage.Korean, "아이슬란드")
	dataIceland.RegisterCapital(xlanguage.Korean, "레이캬비크")
}
