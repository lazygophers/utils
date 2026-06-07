//go:build (lang_ru || lang_all) && (country_africa || country_all || country_cd || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Russian, "ДР Конго")
	dataDrCongo.RegisterOfficialName(xlanguage.Russian, "Демократическая Республика Конго")
	dataDrCongo.RegisterCapital(xlanguage.Russian, "Киншаса")
}
