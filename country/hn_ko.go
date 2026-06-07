//go:build (lang_ko || lang_all) && (country_all || country_americas || country_central_america || country_hn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHonduras.RegisterName(xlanguage.Korean, "온두라스")
	dataHonduras.RegisterOfficialName(xlanguage.Korean, "온두라스 공화국")
	dataHonduras.RegisterCapital(xlanguage.Korean, "테구시갈파")
}
