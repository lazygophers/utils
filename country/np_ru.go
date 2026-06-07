//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Russian, "Непал")
	dataNepal.RegisterOfficialName(xlanguage.Russian, "Федеративная Демократическая Республика Непал")
	dataNepal.RegisterCapital(xlanguage.Russian, "Катманду")
}
