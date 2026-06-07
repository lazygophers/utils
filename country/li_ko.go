//go:build (lang_ko || lang_all) && (country_all || country_europe || country_li || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.Korean, "리히텐슈타인")
	dataLiechtenstein.RegisterOfficialName(xlanguage.Korean, "리히텐슈타인 공국")
	dataLiechtenstein.RegisterCapital(xlanguage.Korean, "파두츠")
}
