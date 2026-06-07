//go:build (lang_ko || lang_all) && (country_all || country_americas || country_aw || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.Korean, "아루바")
	dataAruba.RegisterOfficialName(xlanguage.Korean, "아루바")
	dataAruba.RegisterCapital(xlanguage.Korean, "오라녜스타트")
}
