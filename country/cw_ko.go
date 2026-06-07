//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_cw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.Korean, "퀴라소")
	dataCuracao.RegisterOfficialName(xlanguage.Korean, "퀴라소")
	dataCuracao.RegisterCapital(xlanguage.Korean, "빌렘스타트")
}
