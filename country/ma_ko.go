//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMorocco.RegisterName(xlanguage.Korean, "모로코")
	dataMorocco.RegisterOfficialName(xlanguage.Korean, "모로코 왕국")
	dataMorocco.RegisterCapital(xlanguage.Korean, "라바트")
}
