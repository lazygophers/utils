//go:build (lang_ar || lang_all) && (country_africa || country_all || country_tg || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTogo.RegisterName(xlanguage.Arabic, "توغو")
	dataTogo.RegisterOfficialName(xlanguage.Arabic, "الجمهورية التوغولية")
	dataTogo.RegisterCapital(xlanguage.Arabic, "لومي")
}
