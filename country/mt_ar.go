//go:build (lang_ar || lang_all) && (country_all || country_europe || country_mt || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalta.RegisterName(xlanguage.Arabic, "مالطا")
	dataMalta.RegisterOfficialName(xlanguage.Arabic, "جمهورية مالطا")
	dataMalta.RegisterCapital(xlanguage.Arabic, "فاليتا")
}
