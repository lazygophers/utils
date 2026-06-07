package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Arabic, "تونس")
	dataTunisia.RegisterOfficialName(xlanguage.Arabic, "الجمهورية التونسية")
	dataTunisia.RegisterCapital(xlanguage.Arabic, "تونس")
}
