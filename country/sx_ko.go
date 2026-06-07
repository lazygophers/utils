//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_sx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Korean, "신트마르턴")
	dataSintMaarten.RegisterOfficialName(xlanguage.Korean, "신트마르턴")
	dataSintMaarten.RegisterCapital(xlanguage.Korean, "필립스뷔르흐")
}
