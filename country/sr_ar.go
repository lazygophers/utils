//go:build (lang_ar || lang_all) && (country_all || country_americas || country_south_america || country_sr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Arabic, "سورينام")
	dataSuriname.RegisterOfficialName(xlanguage.Arabic, "جمهورية سورينام")
	dataSuriname.RegisterCapital(xlanguage.Arabic, "باراماريبو")
}
