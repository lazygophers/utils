//go:build (lang_ko || lang_all) && (country_all || country_asia || country_my || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalaysia.RegisterName(xlanguage.Korean, "말레이시아")
	dataMalaysia.RegisterOfficialName(xlanguage.Korean, "말레이시아")
	dataMalaysia.RegisterCapital(xlanguage.Korean, "쿠알라룸푸르")
}
