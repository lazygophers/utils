//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.Korean, "모나코")
	dataMonaco.RegisterOfficialName(xlanguage.Korean, "모나코 공국")
	dataMonaco.RegisterCapital(xlanguage.Korean, "모나코")
}
