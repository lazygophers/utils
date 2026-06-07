//go:build (lang_ar || lang_all) && (country_all || country_asia || country_lk || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSriLanka.RegisterName(xlanguage.Arabic, "سريلانكا")
	dataSriLanka.RegisterOfficialName(xlanguage.Arabic, "جمهورية سريلانكا الديمقراطية الاشتراكية")
	dataSriLanka.RegisterCapital(xlanguage.Arabic, "سري جايواردنابورا كوته")
}
