//go:build (lang_ru || lang_all) && (country_al || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.Russian, "Албания")
	dataAlbania.RegisterOfficialName(xlanguage.Russian, "Республика Албания")
	dataAlbania.RegisterCapital(xlanguage.Russian, "Тирана")
}
