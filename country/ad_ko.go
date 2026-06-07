//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAndorra.RegisterName(xlanguage.Korean, "안도라")
	dataAndorra.RegisterOfficialName(xlanguage.Korean, "안도라 공국")
	dataAndorra.RegisterCapital(xlanguage.Korean, "안도라라베야")
}
