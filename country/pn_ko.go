//go:build (lang_ko || lang_all) && (country_all || country_oceania || country_pn || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.Korean, "핏케언 제도")
	dataPitcairn.RegisterOfficialName(xlanguage.Korean, "핏케언 제도")
	dataPitcairn.RegisterCapital(xlanguage.Korean, "애덤스타운")
}
