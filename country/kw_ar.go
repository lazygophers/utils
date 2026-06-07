package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.Arabic, "الكويت")
	dataKuwait.RegisterOfficialName(xlanguage.Arabic, "دولة الكويت")
	dataKuwait.RegisterCapital(xlanguage.Arabic, "مدينة الكويت")
}
