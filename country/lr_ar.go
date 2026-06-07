//go:build (lang_ar || lang_all) && (country_africa || country_all || country_lr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.Arabic, "ليبيريا")
	dataLiberia.RegisterOfficialName(xlanguage.Arabic, "جمهورية ليبيريا")
	dataLiberia.RegisterCapital(xlanguage.Arabic, "مونروفيا")
}
