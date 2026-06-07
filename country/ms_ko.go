//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_ms)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontserrat.RegisterName(xlanguage.Korean, "몬트세랫")
	dataMontserrat.RegisterOfficialName(xlanguage.Korean, "몬트세랫")
	dataMontserrat.RegisterCapital(xlanguage.Korean, "브레이즈")
}
