//go:build (lang_ko || lang_all) && (country_all || country_americas || country_bo || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBolivia.RegisterName(xlanguage.Korean, "볼리비아")
	dataBolivia.RegisterOfficialName(xlanguage.Korean, "볼리비아 다민족국")
	dataBolivia.RegisterCapital(xlanguage.Korean, "수크레")
}
