//go:build (lang_ru || lang_all) && (country_africa || country_all || country_southern_africa || country_sz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.Russian, "Эсватини")
	dataEswatini.RegisterOfficialName(xlanguage.Russian, "Королевство Эсватини")
	dataEswatini.RegisterCapital(xlanguage.Russian, "Мбабане")
}
