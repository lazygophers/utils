//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_sk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.Arabic, "سلوفاكيا")
	dataSlovakia.RegisterOfficialName(xlanguage.Arabic, "الجمهورية السلوفاكية")
	dataSlovakia.RegisterCapital(xlanguage.Arabic, "براتيسلافا")
}
