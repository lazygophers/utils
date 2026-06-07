//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataQatar.RegisterName(xlanguage.Korean, "카타르")
	dataQatar.RegisterOfficialName(xlanguage.Korean, "카타르국")
	dataQatar.RegisterCapital(xlanguage.Korean, "도하")
}
