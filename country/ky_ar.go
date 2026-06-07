//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_ky)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.Arabic, "جزر كايمان")
	dataCaymanIslands.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر كايمان البريطاني فيما وراء البحار")
	dataCaymanIslands.RegisterCapital(xlanguage.Arabic, "جورج تاون")
}
