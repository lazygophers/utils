//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Russian, "Суринам")
	dataSuriname.RegisterOfficialName(xlanguage.Russian, "Республика Суринам")
	dataSuriname.RegisterCapital(xlanguage.Russian, "Парамарибо")
}
