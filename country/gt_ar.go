//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.Arabic, "غواتيمالا")
	dataGuatemala.RegisterOfficialName(xlanguage.Arabic, "جمهورية غواتيمالا")
	dataGuatemala.RegisterCapital(xlanguage.Arabic, "مدينة غواتيمالا")
}
