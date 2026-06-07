package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLebanon.RegisterName(xlanguage.Arabic, "لبنان")
	dataLebanon.RegisterOfficialName(xlanguage.Arabic, "الجمهورية اللبنانية")
	dataLebanon.RegisterCapital(xlanguage.Arabic, "بيروت")
}
