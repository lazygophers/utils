package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.Arabic, "العراق")
	dataIraq.RegisterOfficialName(xlanguage.Arabic, "جمهورية العراق")
	dataIraq.RegisterCapital(xlanguage.Arabic, "بغداد")
}
