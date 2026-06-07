//go:build (lang_ko || lang_all) && (country_africa || country_all || country_bf || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurkinaFaso.RegisterName(xlanguage.Korean, "부르키나파소")
	dataBurkinaFaso.RegisterOfficialName(xlanguage.Korean, "부르키나파소")
	dataBurkinaFaso.RegisterCapital(xlanguage.Korean, "와가두구")
}
