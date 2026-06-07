//go:build (lang_ja || lang_all) && (country_all || country_asia || country_ge || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGeorgia.RegisterName(xlanguage.Japanese, "ジョージア")
	dataGeorgia.RegisterOfficialName(xlanguage.Japanese, "ジョージア")
	dataGeorgia.RegisterCapital(xlanguage.Japanese, "トビリシ")
}
