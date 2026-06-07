//go:build (lang_ru || lang_all) && (country_all || country_americas || country_bq || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.Russian, "Карибские Нидерланды")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.Russian, "Бонэйр, Синт-Эстатиус и Саба")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.Russian, "Кралендейк")
}
