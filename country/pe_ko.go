//go:build (lang_ko || lang_all) && (country_all || country_americas || country_pe || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.Korean, "페루")
	dataPeru.RegisterOfficialName(xlanguage.Korean, "페루 공화국")
	dataPeru.RegisterCapital(xlanguage.Korean, "리마")
}
