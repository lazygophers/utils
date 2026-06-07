//go:build (lang_ko || lang_all) && (country_all || country_antarctic || country_hm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.Korean, "허드 맥도널드 제도")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.Korean, "허드 맥도널드 제도")
}
