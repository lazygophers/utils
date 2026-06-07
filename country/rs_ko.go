//go:build (lang_ko || lang_all) && (country_all || country_europe || country_rs || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSerbia.RegisterName(xlanguage.Korean, "세르비아")
	dataSerbia.RegisterOfficialName(xlanguage.Korean, "세르비아 공화국")
	dataSerbia.RegisterCapital(xlanguage.Korean, "베오그라드")
}
