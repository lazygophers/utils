//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthSudan.RegisterName(xlanguage.Russian, "Южный Судан")
	dataSouthSudan.RegisterOfficialName(xlanguage.Russian, "Республика Южный Судан")
	dataSouthSudan.RegisterCapital(xlanguage.Russian, "Джуба")
}
