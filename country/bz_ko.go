//go:build (lang_ko || lang_all) && (country_all || country_americas || country_bz || country_central_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Korean, "벨리즈")
	dataBelize.RegisterOfficialName(xlanguage.Korean, "벨리즈")
	dataBelize.RegisterCapital(xlanguage.Korean, "벨모판")
}
