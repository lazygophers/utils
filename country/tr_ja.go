//go:build (lang_ja || lang_all) && (country_all || country_asia || country_tr || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Japanese, "トルコ")
	dataTurkey.RegisterOfficialName(xlanguage.Japanese, "トルコ共和国")
	dataTurkey.RegisterCapital(xlanguage.Japanese, "アンカラ")
}
