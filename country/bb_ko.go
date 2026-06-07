//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.Korean, "바베이도스")
	dataBarbados.RegisterOfficialName(xlanguage.Korean, "바베이도스")
	dataBarbados.RegisterCapital(xlanguage.Korean, "브리지타운")
}
