//go:build (lang_ar || lang_all) && (country_all || country_europe || country_is || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIceland.RegisterName(xlanguage.Arabic, "آيسلندا")
	dataIceland.RegisterOfficialName(xlanguage.Arabic, "جمهورية آيسلندا")
	dataIceland.RegisterCapital(xlanguage.Arabic, "ريكيافيك")
}
