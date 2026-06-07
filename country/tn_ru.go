//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Russian, "Тунис")
	dataTunisia.RegisterOfficialName(xlanguage.Russian, "Тунисская Республика")
	dataTunisia.RegisterCapital(xlanguage.Russian, "Тунис")
}
