//go:build (lang_ko || lang_all) && (country_all || country_asia || country_bn || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Korean, "브루나이")
	dataBrunei.RegisterOfficialName(xlanguage.Korean, "브루나이 다루살람국")
	dataBrunei.RegisterCapital(xlanguage.Korean, "반다르스리브가완")
}
