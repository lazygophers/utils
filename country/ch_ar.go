//go:build (lang_ar || lang_all) && (country_all || country_ch || country_europe || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Arabic, "سويسرا")
	dataSwitzerland.RegisterOfficialName(xlanguage.Arabic, "الكونفدرالية السويسرية")
	dataSwitzerland.RegisterCapital(xlanguage.Arabic, "برن")
}
