//go:build (lang_ar || lang_all) && (country_all || country_asia || country_eastern_asia || country_kp)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthKorea.RegisterName(xlanguage.Arabic, "كوريا الشمالية")
	dataNorthKorea.RegisterOfficialName(xlanguage.Arabic, "جمهورية كوريا الشعبية الديمقراطية")
	dataNorthKorea.RegisterCapital(xlanguage.Arabic, "بيونغ يانغ")
}
