//go:build (lang_ja || lang_all) && (country_all || country_asia || country_om || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataOman.RegisterName(xlanguage.Japanese, "オマーン")
	dataOman.RegisterOfficialName(xlanguage.Japanese, "オマーン国")
	dataOman.RegisterCapital(xlanguage.Japanese, "マスカット")
}
