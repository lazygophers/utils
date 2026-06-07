//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Japanese, "レユニオン")
	dataReunion.RegisterOfficialName(xlanguage.Japanese, "レユニオン")
	dataReunion.RegisterCapital(xlanguage.Japanese, "サン＝ドニ")
}
