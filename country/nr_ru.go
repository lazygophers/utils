//go:build (lang_ru || lang_all) && (country_all || country_micronesia || country_nr || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.Russian, "Науру")
	dataNauru.RegisterOfficialName(xlanguage.Russian, "Республика Науру")
	dataNauru.RegisterCapital(xlanguage.Russian, "Ярен")
}
