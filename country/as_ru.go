//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAmericanSamoa.RegisterName(xlanguage.Russian, "Американское Самоа")
	dataAmericanSamoa.RegisterOfficialName(xlanguage.Russian, "Территория Американское Самоа")
	dataAmericanSamoa.RegisterCapital(xlanguage.Russian, "Паго-Паго")
}
