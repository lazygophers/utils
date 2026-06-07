//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNamibia.RegisterName(xlanguage.Korean, "나미비아")
	dataNamibia.RegisterOfficialName(xlanguage.Korean, "나미비아 공화국")
	dataNamibia.RegisterCapital(xlanguage.Korean, "빈트후크")
}
