//go:build (lang_ko || lang_all) && (country_africa || country_all || country_cv || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.Korean, "카보베르데")
	dataCaboVerde.RegisterOfficialName(xlanguage.Korean, "카보베르데 공화국")
	dataCaboVerde.RegisterCapital(xlanguage.Korean, "프라이아")
}
