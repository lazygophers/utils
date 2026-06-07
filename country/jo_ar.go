//go:build country_all || country_asia || country_jo || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Arabic, "الأردن")
	dataJordan.RegisterOfficialName(xlanguage.Arabic, "المملكة الأردنية الهاشمية")
	dataJordan.RegisterCapital(xlanguage.Arabic, "عمّان")
}
