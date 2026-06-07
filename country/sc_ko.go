//go:build (lang_ko || lang_all) && (country_africa || country_all || country_eastern_africa || country_sc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSeychelles.RegisterName(xlanguage.Korean, "세이셸")
	dataSeychelles.RegisterOfficialName(xlanguage.Korean, "세이셸 공화국")
	dataSeychelles.RegisterCapital(xlanguage.Korean, "빅토리아")
}
