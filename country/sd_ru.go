//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Russian, "Судан")
	dataSudan.RegisterOfficialName(xlanguage.Russian, "Республика Судан")
	dataSudan.RegisterCapital(xlanguage.Russian, "Хартум")
}
