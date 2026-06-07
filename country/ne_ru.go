//go:build (lang_ru || lang_all) && (country_africa || country_all || country_ne || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiger.RegisterName(xlanguage.Russian, "Нигер")
	dataNiger.RegisterOfficialName(xlanguage.Russian, "Республика Нигер")
	dataNiger.RegisterCapital(xlanguage.Russian, "Ниамей")
}
