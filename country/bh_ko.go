//go:build (lang_ko || lang_all) && (country_all || country_asia || country_bh || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.Korean, "바레인")
	dataBahrain.RegisterOfficialName(xlanguage.Korean, "바레인 왕국")
	dataBahrain.RegisterCapital(xlanguage.Korean, "마나마")
}
