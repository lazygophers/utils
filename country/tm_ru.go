//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkmenistan.RegisterName(xlanguage.Russian, "Туркмения")
	dataTurkmenistan.RegisterOfficialName(xlanguage.Russian, "Туркменистан")
	dataTurkmenistan.RegisterCapital(xlanguage.Russian, "Ашхабад")
}
