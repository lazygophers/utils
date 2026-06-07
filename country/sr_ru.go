//go:build (lang_ru || lang_all) && (country_all || country_americas || country_south_america || country_sr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Russian, "Суринам")
	dataSuriname.RegisterOfficialName(xlanguage.Russian, "Республика Суринам")
	dataSuriname.RegisterCapital(xlanguage.Russian, "Парамарибо")
}
