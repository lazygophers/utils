//go:build (lang_ko || lang_all) && (country_all || country_europe || country_northern_europe || country_sj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.Korean, "스발바르 얀마옌")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.Korean, "스발바르 얀마옌")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.Korean, "롱위에아르뷔엔")
}
