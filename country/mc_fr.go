package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.French, "Monaco")
	dataMonaco.RegisterOfficialName(xlanguage.French, "Principauté de Monaco")
	dataMonaco.RegisterCapital(xlanguage.French, "Monaco")
}
