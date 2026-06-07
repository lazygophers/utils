package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.French, "Madagascar")
	dataMadagascar.RegisterOfficialName(xlanguage.French, "République de Madagascar")
	dataMadagascar.RegisterCapital(xlanguage.French, "Antananarivo")
}
