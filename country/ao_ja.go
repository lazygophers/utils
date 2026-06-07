//go:build (lang_ja || lang_all) && (country_africa || country_all || country_ao || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.Japanese, "アンゴラ")
	dataAngola.RegisterOfficialName(xlanguage.Japanese, "アンゴラ共和国")
	dataAngola.RegisterCapital(xlanguage.Japanese, "ルアンダ")
}
