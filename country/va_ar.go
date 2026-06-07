//go:build (lang_ar || lang_all) && (country_all || country_europe || country_southern_europe || country_va)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.Arabic, "الفاتيكان")
	dataVaticanCity.RegisterOfficialName(xlanguage.Arabic, "دولة مدينة الفاتيكان")
	dataVaticanCity.RegisterCapital(xlanguage.Arabic, "مدينة الفاتيكان")
}
