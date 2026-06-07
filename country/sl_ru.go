//go:build (lang_ru || lang_all) && (country_africa || country_all || country_sl || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.Russian, "Сьерра-Леоне")
	dataSierraLeone.RegisterOfficialName(xlanguage.Russian, "Республика Сьерра-Леоне")
	dataSierraLeone.RegisterCapital(xlanguage.Russian, "Фритаун")
}
