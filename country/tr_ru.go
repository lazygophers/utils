//go:build (lang_ru || lang_all) && (country_all || country_asia || country_tr || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTurkey.RegisterName(xlanguage.Russian, "Турция")
	dataTurkey.RegisterOfficialName(xlanguage.Russian, "Турецкая Республика")
	dataTurkey.RegisterCapital(xlanguage.Russian, "Анкара")
}
