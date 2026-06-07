//go:build (lang_ar || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_vn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.Arabic, "فيتنام")
	dataVietnam.RegisterOfficialName(xlanguage.Arabic, "جمهورية فيتنام الاشتراكية")
	dataVietnam.RegisterCapital(xlanguage.Arabic, "هانوي")
}
