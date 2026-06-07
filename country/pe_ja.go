//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.Japanese, "ペルー")
	dataPeru.RegisterOfficialName(xlanguage.Japanese, "ペルー共和国")
	dataPeru.RegisterCapital(xlanguage.Japanese, "リマ")
}
