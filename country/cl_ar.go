//go:build (lang_ar || lang_all) && (country_all || country_americas || country_cl || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChile.RegisterName(xlanguage.Arabic, "تشيلي")
	dataChile.RegisterOfficialName(xlanguage.Arabic, "جمهورية تشيلي")
	dataChile.RegisterCapital(xlanguage.Arabic, "سانتياغو")
}
