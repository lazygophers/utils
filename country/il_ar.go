package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIsrael.RegisterName(xlanguage.Arabic, "إسرائيل")
	dataIsrael.RegisterOfficialName(xlanguage.Arabic, "دولة إسرائيل")
	dataIsrael.RegisterCapital(xlanguage.Arabic, "القدس")
}
