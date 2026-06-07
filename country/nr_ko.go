//go:build (lang_ko || lang_all) && (country_all || country_micronesia || country_nr || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.Korean, "나우루")
	dataNauru.RegisterOfficialName(xlanguage.Korean, "나우루 공화국")
	dataNauru.RegisterCapital(xlanguage.Korean, "야렌")
}
