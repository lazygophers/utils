//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAnguilla.RegisterName(xlanguage.Japanese, "アンギラ")
	dataAnguilla.RegisterOfficialName(xlanguage.Japanese, "アンギラ")
	dataAnguilla.RegisterCapital(xlanguage.Japanese, "ザ・バレー")
}
