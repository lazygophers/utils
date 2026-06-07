//go:build (lang_ar || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Arabic, "زيمبابوي")
	dataZimbabwe.RegisterOfficialName(xlanguage.Arabic, "جمهورية زيمبابوي")
	dataZimbabwe.RegisterCapital(xlanguage.Arabic, "هراري")
}
