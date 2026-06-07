//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVaticanCity.RegisterName(xlanguage.Russian, "Ватикан")
	dataVaticanCity.RegisterOfficialName(xlanguage.Russian, "Государство-город Ватикан")
	dataVaticanCity.RegisterCapital(xlanguage.Russian, "Ватикан")
}
