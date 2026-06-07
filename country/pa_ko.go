//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.Korean, "파나마")
	dataPanama.RegisterOfficialName(xlanguage.Korean, "파나마 공화국")
	dataPanama.RegisterCapital(xlanguage.Korean, "파나마시티")
}
