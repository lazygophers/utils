//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.Russian, "Китай")
	dataChina.RegisterOfficialName(xlanguage.Russian, "Китайская Народная Республика")
	dataChina.RegisterCapital(xlanguage.Russian, "Пекин")
}
