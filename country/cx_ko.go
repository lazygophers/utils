//go:build (lang_ko || lang_all) && (country_all || country_australia_and_new_zealand || country_cx || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChristmasIsland.RegisterName(xlanguage.Korean, "크리스마스섬")
	dataChristmasIsland.RegisterOfficialName(xlanguage.Korean, "크리스마스섬")
	dataChristmasIsland.RegisterCapital(xlanguage.Korean, "플라잉피시코브")
}
