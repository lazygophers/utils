//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAngola.RegisterName(xlanguage.Russian, "Ангола")
	dataAngola.RegisterOfficialName(xlanguage.Russian, "Республика Ангола")
	dataAngola.RegisterCapital(xlanguage.Russian, "Луанда")
}
