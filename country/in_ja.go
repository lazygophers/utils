//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIndia.RegisterName(xlanguage.Japanese, "インド")
	dataIndia.RegisterOfficialName(xlanguage.Japanese, "インド共和国")
	dataIndia.RegisterCapital(xlanguage.Japanese, "ニューデリー")
}
