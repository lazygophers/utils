//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_vc)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.Korean, "세인트빈센트 그레나딘")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.Korean, "세인트빈센트 그레나딘")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.Korean, "킹스타운")
}
