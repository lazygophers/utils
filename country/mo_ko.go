//go:build (lang_ko || lang_all) && (country_all || country_asia || country_eastern_asia || country_mo)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.Korean, "마카오")
	dataMacao.RegisterOfficialName(xlanguage.Korean, "중화인민공화국 마카오 특별행정구")
	dataMacao.RegisterCapital(xlanguage.Korean, "마카오")
}
