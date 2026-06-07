//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.Japanese, "アゼルバイジャン")
	dataAzerbaijan.RegisterOfficialName(xlanguage.Japanese, "アゼルバイジャン共和国")
	dataAzerbaijan.RegisterCapital(xlanguage.Japanese, "バクー")
}
