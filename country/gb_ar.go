//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.Arabic, "المملكة المتحدة")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.Arabic, "المملكة المتحدة لبريطانيا العظمى وأيرلندا الشمالية")
	dataUnitedKingdom.RegisterCapital(xlanguage.Arabic, "لندن")
}
