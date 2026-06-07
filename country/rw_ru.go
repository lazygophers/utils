//go:build (lang_ru || lang_all) && (country_africa || country_all || country_eastern_africa || country_rw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.Russian, "Руанда")
	dataRwanda.RegisterOfficialName(xlanguage.Russian, "Республика Руанда")
	dataRwanda.RegisterCapital(xlanguage.Russian, "Кигали")
}
