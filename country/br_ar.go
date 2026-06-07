//go:build (lang_ar || lang_all) && (country_all || country_americas || country_br || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.Arabic, "البرازيل")
	dataBrazil.RegisterOfficialName(xlanguage.Arabic, "جمهورية البرازيل الاتحادية")
	dataBrazil.RegisterCapital(xlanguage.Arabic, "برازيليا")
}
