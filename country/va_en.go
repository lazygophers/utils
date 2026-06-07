package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.English, "Vatican City")
	dataVaticanCity.RegisterOfficialName(xlanguage.English, "Vatican City State")
	dataVaticanCity.RegisterCapital(xlanguage.English, "Vatican City")
}
