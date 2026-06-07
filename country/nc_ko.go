//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.Korean, "누벨칼레도니")
	dataNewCaledonia.RegisterOfficialName(xlanguage.Korean, "누벨칼레도니")
	dataNewCaledonia.RegisterCapital(xlanguage.Korean, "누메아")
}
