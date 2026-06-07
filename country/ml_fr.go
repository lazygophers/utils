package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMali.RegisterName(xlanguage.French, "Mali")
	dataMali.RegisterOfficialName(xlanguage.French, "République du Mali")
	dataMali.RegisterCapital(xlanguage.French, "Bamako")
}
