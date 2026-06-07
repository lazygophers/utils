//go:build (lang_ar || lang_all) && (country_africa || country_all || country_bj || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBenin.RegisterName(xlanguage.Arabic, "بنين")
	dataBenin.RegisterOfficialName(xlanguage.Arabic, "جمهورية بنين")
	dataBenin.RegisterCapital(xlanguage.Arabic, "بورتو نوفو")
}
