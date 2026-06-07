//go:build (lang_ar || lang_all) && (country_all || country_europe || country_ie || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIreland.RegisterName(xlanguage.Arabic, "أيرلندا")
	dataIreland.RegisterOfficialName(xlanguage.Arabic, "جمهورية أيرلندا")
	dataIreland.RegisterCapital(xlanguage.Arabic, "دبلن")
}
