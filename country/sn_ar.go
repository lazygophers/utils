//go:build (lang_ar || lang_all) && (country_africa || country_all || country_sn || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.Arabic, "السنغال")
	dataSenegal.RegisterOfficialName(xlanguage.Arabic, "جمهورية السنغال")
	dataSenegal.RegisterCapital(xlanguage.Arabic, "داكار")
}
