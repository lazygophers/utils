//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Korean, "몽골")
	dataMongolia.RegisterOfficialName(xlanguage.Korean, "몽골")
	dataMongolia.RegisterCapital(xlanguage.Korean, "울란바토르")
}
