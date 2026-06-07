package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Arabic, "الأردن")
	dataJordan.RegisterOfficialName(xlanguage.Arabic, "المملكة الأردنية الهاشمية")
	dataJordan.RegisterCapital(xlanguage.Arabic, "عمّان")
}
