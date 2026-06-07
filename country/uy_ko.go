//go:build (lang_ko || lang_all) && (country_all || country_americas || country_south_america || country_uy)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUruguay.RegisterName(xlanguage.Korean, "우루과이")
	dataUruguay.RegisterOfficialName(xlanguage.Korean, "우루과이 동방 공화국")
	dataUruguay.RegisterCapital(xlanguage.Korean, "몬테비데오")
}
