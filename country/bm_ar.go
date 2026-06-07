//go:build (lang_ar || lang_all) && (country_all || country_americas || country_bm || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.Arabic, "برمودا")
	dataBermuda.RegisterOfficialName(xlanguage.Arabic, "إقليم برمودا البريطاني فيما وراء البحار")
	dataBermuda.RegisterCapital(xlanguage.Arabic, "هاميلتون")
}
