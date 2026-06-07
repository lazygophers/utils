//go:build (lang_ru || lang_all) && (country_all || country_asia || country_eastern_asia || country_mo)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.Russian, "Макао")
	dataMacao.RegisterOfficialName(xlanguage.Russian, "Специальный административный район Аомынь Китайской Народной Республики")
	dataMacao.RegisterCapital(xlanguage.Russian, "Макао")
}
