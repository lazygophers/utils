//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArgentina.RegisterName(xlanguage.Japanese, "アルゼンチン")
	dataArgentina.RegisterOfficialName(xlanguage.Japanese, "アルゼンチン共和国")
	dataArgentina.RegisterCapital(xlanguage.Japanese, "ブエノスアイレス")
}
