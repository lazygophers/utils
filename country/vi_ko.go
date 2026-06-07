//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.Korean, "미국령 버진아일랜드")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.Korean, "미국령 버진아일랜드")
	dataUsVirginIslands.RegisterCapital(xlanguage.Korean, "샬럿아말리에")
}
