//go:build (lang_ar || lang_all) && (country_all || country_eastern_europe || country_europe || country_md)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.Arabic, "مولدوفا")
	dataMoldova.RegisterOfficialName(xlanguage.Arabic, "جمهورية مولدوفا")
	dataMoldova.RegisterCapital(xlanguage.Arabic, "كيشيناو")
}
