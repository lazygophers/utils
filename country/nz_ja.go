//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Japanese, "ニュージーランド")
	dataNewZealand.RegisterOfficialName(xlanguage.Japanese, "ニュージーランド")
	dataNewZealand.RegisterCapital(xlanguage.Japanese, "ウェリントン")
}
