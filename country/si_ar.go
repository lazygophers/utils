//go:build (lang_ar || lang_all) && (country_all || country_europe || country_si || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovenia.RegisterName(xlanguage.Arabic, "سلوفينيا")
	dataSlovenia.RegisterOfficialName(xlanguage.Arabic, "جمهورية سلوفينيا")
	dataSlovenia.RegisterCapital(xlanguage.Arabic, "ليوبليانا")
}
