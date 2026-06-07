//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.Russian, "Германия")
	dataGermany.RegisterOfficialName(xlanguage.Russian, "Федеративная Республика Германия")
	dataGermany.RegisterCapital(xlanguage.Russian, "Берлин")
}
