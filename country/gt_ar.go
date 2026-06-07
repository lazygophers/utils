//go:build (lang_ar || lang_all) && (country_all || country_americas || country_central_america || country_gt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.Arabic, "غواتيمالا")
	dataGuatemala.RegisterOfficialName(xlanguage.Arabic, "جمهورية غواتيمالا")
	dataGuatemala.RegisterCapital(xlanguage.Arabic, "مدينة غواتيمالا")
}
