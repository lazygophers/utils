//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataIraq.RegisterName(xlanguage.Russian, "Ирак")
	dataIraq.RegisterOfficialName(xlanguage.Russian, "Республика Ирак")
	dataIraq.RegisterCapital(xlanguage.Russian, "Багдад")
}
